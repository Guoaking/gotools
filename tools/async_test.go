package tools

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
@description
@date: 08/08 11:36
@author Gk
**/

// TestTas 测试任务
func TestTas(t *testing.T) {

	//rand.Seed(time.Now().UnixNano())

	pool := NewAsyncTaskPool(2, 1, RetryPolicy{Retries: 3, WaitTime: time.Second})
	go func() {
		for result := range pool.GetResultChan() {
			fmt.Sprintf("result: %v, err: %v\n", result.Result, result.Err)
		}
	}()

	for i := 0; i < 5; i++ {
		//task := NewAsyncTask(func(ctx context.Context, args interface{}) (interface{}, error) {
		//	num := args.(int)
		//	fmt.Printf("task %d start\n", num)
		//	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		//	fmt.Printf("task %d end\n", num)
		//	return num * 2, nil
		//}, i, time.Second, 2, time.Second*3)
		pool.AddTaskWithPriority(func(ctx context.Context, args interface{}) (interface{}, error) {
			num := args.(int)

			duration := time.Duration(rand.Intn(10)) * time.Second
			fmt.Printf("task %d start %v<<< %v\n", num, time.Now(), duration)
			time.Sleep(duration)
			//fmt.Printf("task %d end\n", num)
			return num * 2, nil
		}, i, 1, 2, time.Second*3)
	}

	//task1 := NewAsyncTask((&TaskFunc{TaskInput{"Alice", 25}}).Run, 1, 1, time.Second*3, 2, time.Second)
	//task2 := NewAsyncTask((&TaskFunc{TaskInput{"Bob", 30}}).Run, 2, 2, time.Second*3, 1, time.Second)
	//task3 := NewAsyncTask((&TaskFunc2{TaskInput{"Bob", 30}}).Run, 2, 2, time.Second*3, 1, time.Second)

	//
	//pool.AddTask(task1)
	//pool.AddTask(task2)
	//pool.AddTaskWithPriority(task3)

	//fmt.Printf("output:len %v\n", pool.priorityQueue.Len())
	pool.Start()
	//fmt.Println("1111")
	//for result := range pool.GetResultChan() {
	//	fmt.Printf("result: %v, err: %v\n", result.result, result.err)
	//}

	//for result := range pool.GetResultChan() {
	//	output := result.result.(TaskOutput)
	//	fmt.Printf("result1: %s, result2: %d, err: %v\n", output.Result1, output.Result2, result.err)
	//}

	// 等待一段时间后关闭任务池
	//time.Sleep(time.Second * 10)
	//pool.Close()

}
