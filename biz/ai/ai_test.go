package ai

import (
	"context"
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
	getJson()

}

func GetMetaJson(dir string) []string {
	var res []string
	One1(&res, dir, func(entry os.DirEntry) bool {
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

func One1(res *[]string, dir string, f func(entry os.DirEntry) bool) {
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
		One1(res, d1, f)
	}
}

func getJson() {
	dir := "/Users/bytedance/Downloads/图片"
	background := context.Background()
	files := GetMetaJson(dir)
	for _, file := range files {
		readJson := ReadJson(file)
		base, desc := GetCNInfo(readJson)

		if base == "" || desc == "" {
			continue
		}
		data := Msg{
			BaseName: base,
			DescALl:  desc,
		}

		marshal, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("output: %v\n", err)
			return
		}
		fmt.Printf("output: %s\n", marshal)

		one := One(background, string(marshal))

		var D2 ali.TaD
		err = json.Unmarshal([]byte(one), &D2)
		if err != nil {
			fmt.Printf("output: %v\n", err)
			return
		}

		product, ok := readJson["zh-cn"]
		if !ok {
			return
		}

		product.Title = D2.Title
		product.Desc = D2.Desc
		readJson["zh-cn"] = product

		marshal, err3 := json.Marshal(readJson)
		if err3 != nil {
			fmt.Printf("output: %v\n", err3)
			return
		}

		name := file + ".txt"
		fmt.Printf("output:name %v\n", name)
		err3 = os.WriteFile(name, marshal, os.ModePerm)
		if err3 != nil {
			fmt.Printf("output: %v\n", err3)
			return
		}

		break
	}
}

type Msg struct {
	BaseName string `json:"base_name"`
	DescALl  string `json:"desc_all"`
}

func GetCNInfo(D ali.I18nProduct) (string, string) {
	product, ok := D["zh-cn"]
	if !ok || len(product.DescAll) == 0 || product.Title != "" || product.Desc != "" {
		return "", ""
	}
	return product.BaseName, strings.Join(product.DescAll, ",")
}

func ReadJson(filename string) ali.I18nProduct {

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file error: %v\n", err)
		return nil
	}

	var D ali.I18nProduct
	err = json.Unmarshal(file, &D)
	if err != nil {
		fmt.Printf("json unmarshal error: %v\n", err)
		return nil
	}

	return D
}

func Test2(t *testing.T) {
	one := `{
    "status": "success",
    "data": {
        "title": "高清航拍玩具无人机 稳定悬停",
        "desc": "升级无刷电机稳定飞行，高清航拍远近皆清晰，带屏遥控操作便捷，光流定位悬停更安全。"
    }
}`
	var D2 ali.Ai
	err := json.Unmarshal([]byte(one), &D2)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return
	}

	fmt.Printf("output:D2 %v\n", D2)
}
