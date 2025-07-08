package ali

import (
	"fmt"
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
