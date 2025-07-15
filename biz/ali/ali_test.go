package ali

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"Guoaking/gotools/tools"
)

/**
@description
@date: 07/07 14:47
@author Gk
**/

func TestAli(t *testing.T) {

	dirAny, err := tools.ListFilesInCurrentDirAny("/Users/bytedance/Documents/图片/莆田鞋子")
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	for _, sourcePath := range dirAny {
		fmt.Printf("output: %v\n", sourcePath)

		err := TransImages(sourcePath)
		if err != nil {
			fmt.Printf("output: %v\n", err)
			return
		}

	}

	//DownImage("https://cdn.translate.alibaba.com/r/1e55a01512a24f85b67c2317cff59fae.jpg")
}

func TestAli2(t *testing.T) {
	TransImage("/Users/bytedance/Documents/图片/蓝牙耳机/5/IMG_0070.JPG")
}

func TestStd(t *testing.T) {
	// 当前目录下 创建一个 zh-cn的文件夹, 把当前目录下所有的图片, 移动到这个目录下
	//All()
	//ReNameTo("/Users/bytedance/Downloads/图片/磁吸充电宝2", "zh-tw")
	// 文件数量不超过10个
	// 翻译时如果没有识别到任何文字的纯图, 做一个标记
	fmt.Printf("output: %v\n", tools.ProcessDirs("/Users/bytedance/Downloads/图片/"))
}
func MvTo(dirPath, NewDirName string) {
	name := NewDirName
	dirAny, err := tools.ListFilesInCurrentDirAny(dirPath)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	for _, sourcePath := range dirAny {
		new1 := path.Join(sourcePath, name)
		tools.MoveFiles(sourcePath+"/*.JPG", new1)
		tools.MoveFiles(sourcePath+"/*.PNG", new1)
	}
}

func ReNameTo(dirPath, NewDirName string) {
	//name := NewDirName
	dirAny, err := tools.ListFilesInCurrentDirAny(dirPath)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	for _, sourcePath := range dirAny {
		//fmt.Printf("output: %v\n", sourcePath)

		dirAny2, err2 := tools.ListFilesInCurrentDirAny(sourcePath)
		if err2 != nil {
			fmt.Printf("output: %v\n", err)
			return
		}
		for _, sourcePath2 := range dirAny2 {
			new2 := path.Join(sourcePath, NewDirName)
			base := path.Base(sourcePath2)

			if base == "f" {
				fmt.Printf("output: %v\n", tools.RenameDirectory(sourcePath2, new2))
				//fmt.Printf("output: %v\n", os.RemoveAll(new2))
			}

			break
		}

	}
}

func All() {
	name := "zh-cn"
	dirAny, err := tools.ListFilesInCurrentDirAny("/Users/bytedance/Downloads/图片/")
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	for _, sourcePath := range dirAny {
		MvTo(sourcePath, name)
	}
}

func TestOcr(t *testing.T) {
	//给到 一组数据

	//OcrImage("/Users/bytedance/Downloads/图片/磁吸充电宝/2/zh-cn/IMG_0270.JPG")
	// 忽略无意义的字母, 字符, 单词, 符号,
	// 尝试提取,总结出抓人眼球的 内容
	// 标题(控制在20个字以内)
	// 描述(控制在50个字以内), 不需要精确的数字内容(500, 2000之类的, 只要模糊的文字)

	dir := "/Users/bytedance/Downloads/图片"
	//OneTypeManyProduct(dir, "zh-cn", []string{"jpg", "png"})
	ManyTypeManyProduct(dir, "zh-cn", []string{"jpg", "png"})
}

func ManyTypeManyProduct(dir, lang string, f []string) {

	filter, err := tools.ListFilesInCurrentDirFilter2(dir, func(entry os.DirEntry) bool {
		return entry.IsDir()
	})

	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	for _, p := range filter {
		fmt.Printf("output: %v\n", p)
		OneTypeManyProduct(p, lang, f)
	}
}

func OneTypeManyProduct(dir, lang string, f []string) {

	filter, err := tools.ListFilesInCurrentDirFilter2(dir, func(entry os.DirEntry) bool {
		return entry.IsDir()
	})

	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	for _, p := range filter {
		OCROneProduct(p, lang, f)
	}

}

func OCROneProduct(dir, lang string, f []string) {
	meta := path.Join(dir, "meta.json")
	var D I18nProduct

	_, err3 := os.Stat(meta)
	if err3 == nil {
		file, err2 := os.ReadFile(meta)
		if err2 != nil {
			fmt.Printf("output: %v\n", err2)
			return
		}
		fmt.Printf("output: %s\n", file)
		err2 = json.Unmarshal(file, &D)
		if err2 != nil {
			fmt.Printf("output: %v\n", err2)
			return
		}
	}

	var descall []string
	if D == nil {
		D = I18nProduct{}
	}

	dir = path.Join(dir, lang)
	filter, err := tools.ListFilesInCurrentDirFilter2(dir, func(entry os.DirEntry) bool {
		for _, s := range f {
			if strings.Contains(strings.ToLower(entry.Name()), s) {
				return true
			}
		}
		return false
	})
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	for _, p := range filter {
		res, err := OcrImage(p)
		if err != nil {
			fmt.Printf("output: %v\n", err)
			return
		}
		descall = append(descall, res)
	}

	baseName := path.Base(path.Dir(path.Dir(dir)))

	product, ok := D["zh-cn"]
	if ok {
		if product.BaseName == "" {
			product.BaseName = baseName
		}

		product.DescAll = append(product.DescAll, descall...)
		D["zh-cn"] = product
	} else {
		D["zh-cn"] = LocalProduct{
			DescAll:  descall,
			BaseName: baseName,
		}
	}

	marshal, err3 := json.Marshal(D)
	if err3 != nil {
		fmt.Printf("output: %v\n", err3)
		return
	}
	err3 = os.WriteFile(meta, marshal, os.ModePerm)
	if err3 != nil {
		fmt.Printf("output: %v\n", err3)
		return
	}

}

func TestTrans(t *testing.T) {
	// 读取json的参数

	res, err := TransTxt("无人机", "zh-tw")
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	fmt.Printf("output: %v\n", res)

}
