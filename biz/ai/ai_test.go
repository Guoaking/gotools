package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
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
	step1AI()
	//step2Check()
	step3TransTo()
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
			*res = append(*res, d1)
			continue
		}
		One1(res, d1, f)
	}
}

func step1AI() {
	dir := "/Users/bytedance/Downloads/图片"
	background := context.Background()
	files := GetMetaJson(dir)
	for _, file := range files {
		readJson := ReadJson(file)

		product, ok := readJson["zh-cn"]

		if !ok || len(product.DescAll) == 0 || product.Title != "" || product.Desc != "" {
			continue
		}

		data := Msg{
			BaseName: product.BaseName,
			DescALl:  strings.Join(product.DescAll, ","),
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
			fmt.Printf("output:获取dp的结果失败:  %v\n", err)
			return
		}

		product.Title = D2.Title
		product.Desc = D2.Desc
		product.Dir = path.Join(path.Dir(file), "zh-cn")
		readJson["zh-cn"] = product

		marshal, err3 := json.Marshal(readJson)
		if err3 != nil {
			fmt.Printf("output: %v\n", err3)
			return
		}

		name := file
		fmt.Printf("output:name %v\n", name)
		err3 = os.WriteFile(name, marshal, os.ModePerm)
		if err3 != nil {
			fmt.Printf("output: %v\n", err3)
			return
		}

	}
}

func step2Check() {
	dir := "/Users/bytedance/Downloads/图片"
	files := GetMetaJson(dir)
	for _, file := range files {
		readJson := ReadJson(file)

		product, ok := readJson["zh-cn"]

		if !ok {
			continue
		}

		if product.Title == "" || product.Desc == "" {

			fmt.Printf("output: %v\n", file)
		}
	}
}

type Msg struct {
	BaseName string `json:"base_name"`
	DescALl  string `json:"desc_all"`
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

func step3TransTo() {
	dir := "/Users/bytedance/Downloads/图片"
	files := GetMetaJson(dir)

	targetLang := "zh-tw"

	for _, file := range files {
		readJson := ReadJson(file)

		targetprod, ok := readJson[targetLang]
		if ok && targetprod.Title != "" && targetprod.Desc != "" {
			continue
		}

		product, ok := readJson["zh-cn"]
		if !ok || product.Title == "" || product.Desc == "" {
			continue
		}

		title, err := ali.TransTxt(product.Title, targetLang)
		if err != nil {
			fmt.Printf("output1: %v\n", err)
			return
		}
		desc, err := ali.TransTxt(product.Desc, targetLang)
		if err != nil {
			fmt.Printf("output2: %v\n", err)
			return
		}

		//拿到结果 组装并保存
		readJson[targetLang] = ali.LocalProduct{
			TaD: ali.TaD{
				Title: title,
				Desc:  desc,
			},
			Dir:   path.Join(path.Dir(file), targetLang),
			Price: "",
			Cate:  product.Cate,
		}

		marshal, err3 := json.Marshal(readJson)
		if err3 != nil {
			fmt.Printf("output: %v\n", err3)
			return
		}

		name := file
		fmt.Printf("output:name %v\n", name)
		err3 = os.WriteFile(name, marshal, os.ModePerm)
		if err3 != nil {
			fmt.Printf("output: %v\n", err3)
			return
		}

	}
}
