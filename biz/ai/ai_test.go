package ai

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"Guoaking/gotools/biz/ali"
	"Guoaking/gotools/tools"
)

/**
@description
@date: 07/13 19:44
@author Gk
**/

func TestA(t *testing.T) {
	//o1()
	getJson()
}

func GetMetaJson(dir string) []string {
	var res []string
	One(&res, dir, func(entry os.DirEntry) bool {
		if strings.Contains(strings.ToLower(entry.Name()), "meta.json") {
			return true
		}

		if strings.Contains(strings.ToLower(entry.Name()), "zh-tw") {
			return false
		}
		if strings.Contains(strings.ToLower(entry.Name()), "zh-cn") {
			return false
		}

		return entry.IsDir()
	})
	return res
}

func One(res *[]string, dir string, f func(entry os.DirEntry) bool) {
	dirs, err := tools.ListFilesInCurrentDirFilter2(dir, f)

	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	for _, d1 := range dirs {
		if strings.Contains(strings.ToLower(d1), "meta.json") {
			//fmt.Printf("output: %v, %v \n", i, d1)
			*res = append(*res, d1)
			continue
		}
		One(res, d1, f)
	}
}

func getJson() {
	dir := "/Users/bytedance/Downloads/图片"
	files := GetMetaJson(dir)
	for _, file := range files {
		ReadJson(file)
	}
}

func ReadJson(filename string) {

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file error: %v\n", err)
		return
	}

	var D ali.I18nProduct
	err = json.Unmarshal(file, &D)
	if err != nil {
		fmt.Printf("json unmarshal error: %v\n", err)
		return
	}

	product, ok := D["zh-cn"]
	if !ok || len(product.DescAll) == 0 || product.Title != "" || product.Desc != "" {
		return
	}

	//product.BaseName
	//product.DescAll

	//func ()
	//return title, desc
}
