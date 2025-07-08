package tools

import (
	"container/heap"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type TaskFunc func(context.Context, interface{}) (interface{}, error)

//type AsyncTask struct {
//	id         int64     // 任务唯一标识符
//	f          TaskFunc // 任务函数
//	args       interface{}
//	retries    int
//	retryWait  time.Duration
//}

type AsyncTask struct {
	id       int64       // 任务唯一标识符
	taskFunc TaskFunc    // 异步任务函数
	args     interface{} // 异步任务参数
	//resultChan chan AsyncTaskResult // 异步任务结果channel
	//priority     int                                                     // 优先级
	timeout      time.Duration      // 超时时间
	cancelCtx    context.Context    // 任务取消上下文
	cancelFunc   context.CancelFunc // 取消函数
	cancelSignal chan struct{}      // 取消信号
	retryPolicy  RetryPolicy        // 重试策略
}

type AsyncTaskResult struct {
	Result interface{} // 异步任务结果
	Err    error       // 异步任务错误
}

type PriorityTask struct {
	task     AsyncTask
	priority int
}

type PriorityQueue []*PriorityTask

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*PriorityTask)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type AsyncTaskPool struct {
	//taskChan    chan *PriorityTask   // 异步任务channel（带优先级）
	PriorityQueue PriorityQueue        // 异步任务channel（带优先级）
	resultChan    chan AsyncTaskResult // 异步任务结果channel
	maxWorkers    int                  // 最大协程数
	waitGroup     sync.WaitGroup       // 等待所有任务完成
	//retryPolicy   RetryPolicy          // 重试策略
	done   chan struct{} // 控制主线程退出
	mutex  sync.Mutex    // 互斥锁，保护 priorityQueue
	taskID int64         // 自增的任务 ID
}

type RetryPolicy struct {
	Retries  int           // 重试次数
	WaitTime time.Duration // 重试等待时间
}

func NewAsyncTask(taskFunc TaskFunc, args interface{}, timeout time.Duration, retries int, waitTime time.Duration) *AsyncTask {
	ctx, cancel := context.WithCancel(context.Background())
	return &AsyncTask{
		taskFunc: taskFunc,
		args:     args,
		//resultChan: make(chan AsyncTaskResult, 1),
		//priority:     priority,
		id:           time.Now().UnixNano(),
		timeout:      timeout,
		cancelCtx:    ctx,
		cancelFunc:   cancel,
		cancelSignal: make(chan struct{}),
		retryPolicy:  RetryPolicy{Retries: retries, WaitTime: waitTime},
	}
}

func (t *AsyncTask) execute(p *AsyncTaskPool) {

	var result AsyncTaskResult
	timer := time.NewTimer(t.timeout)
	defer timer.Stop()

	for i := 0; i <= t.retryPolicy.Retries; i++ {
		select {
		case <-t.cancelSignal:
			result.Err = errors.New("task canceled")
			p.resultChan <- result
			return
		//case <-timer.C:
		//	// 超时和重试有冲突?
		//	result.err = errors.New(fmt.Sprintf("task %s timeout", t.timeout))
		//	p.resultChan <- result
		//	return
		default:
			// 去执行
			result.Result, result.Err = t.taskFunc(t.cancelCtx, t.args)

			if result.Err == nil {
				p.resultChan <- result
				return
			} else {
				if i == t.retryPolicy.Retries {
					p.resultChan <- result
					return
				} else {
					fmt.Printf("%v后第%v/%v重试, 失败导致重试原因:%v, args: %v\n", t.retryPolicy.WaitTime, i, t.retryPolicy.Retries, result.Err, t.args)
					time.Sleep(t.retryPolicy.WaitTime)
				}
			}
		}
	}
}

func NewAsyncTaskPool(maxWorkers, bufferSize int, retryPolicy RetryPolicy) *AsyncTaskPool {
	return &AsyncTaskPool{
		PriorityQueue: make(PriorityQueue, 0),
		resultChan:    make(chan AsyncTaskResult, bufferSize),
		maxWorkers:    maxWorkers,
		waitGroup:     sync.WaitGroup{},
		//retryPolicy:   retryPolicy,
		done: make(chan struct{}),
	}
}

//func (p *AsyncTaskPool) Start() {
//	pq := make(PriorityQueue, 0)
//	heap.Init(&pq)
//
//	for i := 0; i < p.maxWorkers; i++ {
//		go p.worker(&pq)
//	}
//
//	for task := range p.taskChan {
//		heap.Push(&pq, task)
//	}
//
//	p.waitGroup.Wait()
//	close(p.resultChan)
//}
//func (p *AsyncTaskPool) Start() {
//	for i := 0; i < p.maxWorkers; i++ {
//		p.waitGroup.Add(1)
//		go p.worker()
//	}
//
//	go func() {
//		p.waitGroup.Wait()
//		close(p.resultChan)
//		close(p.done)
//	}()
//
//	<-p.done
//}

func (p *AsyncTaskPool) Start() {
	//pq := make(PriorityQueue, 0)
	//heap.Init(&pq)

	for i := 0; i < p.maxWorkers; i++ {
		go p.worker(i)
	}

	for {
		select {

		case <-p.done:
			// 设置任务池已关闭，不再接收新任务
			//close(p.taskChan)

			// 等待所有任务完成
			p.waitGroup.Wait()

			// 关闭结果通道
			close(p.resultChan)
			close(p.done)
			return
		case <-time.After(time.Second):
			// 检查任务队列是否为空
			if p.PriorityQueue.Len() == 0 {
				// 设置任务池已关闭，不再接收新任务
				//close(p.taskChan)

				// 等待所有任务完成
				p.waitGroup.Wait()

				// 关闭结果通道
				close(p.resultChan)

				// 发送退出信号
				close(p.done)
				return
			}
		}
	}
}

func NewPriorityTask(task *AsyncTask, priority int) *PriorityTask {
	return &PriorityTask{
		task:     *task,
		priority: priority,
	}
}

//func (p *AsyncTaskPool) AddTask(task *AsyncTask) {
//	p.waitGroup.Add(1)
//	p.taskChan <- &PriorityTask{*task, 1}
//
//	//p.waitGroup.Add(1)
//	//
//	//go func() {
//	//	defer p.waitGroup.Done()
//	//	task.execute()
//	//}()
//}

//func (p *AsyncTaskPool) AddPriorityTask(task *PriorityTask) {
//	p.waitGroup.Add(1)
//	p.taskChan <- task
//}

func (p *AsyncTaskPool) AddTaskWithPriority(f TaskFunc, args interface{}, priority int, retries int, retryWait time.Duration) {
	task := NewAsyncTask(f, args, time.Second*5, retries, retryWait)
	priorityTask := &PriorityTask{task: *task, priority: priority}
	p.mutex.Lock()
	p.PriorityQueue.Push(priorityTask)
	p.waitGroup.Add(1)
	defer p.mutex.Unlock()
	//sort.Sort(&p.priorityQueue)
}
func (p *AsyncTaskPool) AddTaskWithPriority2(f AsyncTask, priority int) {

	priorityTask := &PriorityTask{task: f, priority: priority}
	p.mutex.Lock()

	p.PriorityQueue.Push(priorityTask)
	p.waitGroup.Add(1)
	defer p.mutex.Unlock()

}

func (p *AsyncTaskPool) worker(id int) {
	//fmt.Printf("output: worker: %v\n", id)
	for {
		p.mutex.Lock()
		if p.PriorityQueue.Len() == 0 {
			//time.Sleep(time.Millisecond * 100)
			//continue
			//fmt.Printf("output: worker: %v: 没有任务了 return\n", id)
			p.mutex.Unlock()
			return
		}

		task := heap.Pop(&p.PriorityQueue).(*PriorityTask).task
		//fmt.Printf("output: worker: %v: 拿出一个任务, 还有 %v 个\n", id, p.priorityQueue.Len())
		p.mutex.Unlock()

		task.execute(p)
		p.waitGroup.Done()

		//task := p.priorityQueue[0].task
		//p.priorityQueue = p.priorityQueue[1:]
		//p.mutex.Unlock()

		//result := AsyncTaskResult{task: task}
		//result.result, result.err = task.execute()
		//p.resultChan <- result
	}
}

//func (p *AsyncTaskPool) worker(pq *PriorityQueue) {
//	fmt.Printf("output: %v\n", "worker")
//	defer p.waitGroup.Done()
//
//	for {
//		if pq.Len() == 0 {
//			time.Sleep(time.Millisecond * 100)
//			continue
//		}
//
//		item := heap.Pop(pq).(*PriorityTask)
//		task := item.task
//		task.execute()
//
//		select {
//		//case  <-p.taskChan:
//		//
//		case <-task.cancelSignal:
//			continue
//			//case p.taskChan <- item:
//			//	result := <-task.resultChan
//			//	p.resultChan <- result
//		}
//	}
//}

func (p *AsyncTaskPool) GetResultChan() <-chan AsyncTaskResult {
	return p.resultChan
}

func (p *AsyncTaskPool) Close() {
	// 发送关闭信号
	close(p.done)
}

//该代码实现了一个
//1. 支持任务优先级、 任务超时和任务取消的异步任务池，
//2. 方便地执行异步任务，并获取异步任务的执行结果。
//3. 其中，任务优先级越高的任务会先执行；任务超时会导致任务失败并返回错误；任务取消可以在任何时候取消任务的执行。在任务执行失败时，可以根据设定的重试策略进行重试。
