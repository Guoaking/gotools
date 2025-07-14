package ali

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"Guoaking/gotools/tools"

	alimt20181012 "github.com/alibabacloud-go/alimt-20181012/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ocr_api20210707 "github.com/alibabacloud-go/ocr-api-20210707/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"
)

/**
@description
@date: 07/07 14:18
@author Gk
**/

// This file is auto-generated, don't edit it. Thanks.
// Description:
// 使用凭据初始化账号Client
// @return Client
// @throws Exception

func getBasicConfig() *openapi.Config {
	secret := new(credential.Config).
		SetType("access_key").
		// 从环境变量中获取AccessKey Id。
		SetAccessKeyId(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")).
		// 从环境变量中获取AccessKey Secret
		SetAccessKeySecret(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"))
	cred, _err := credential.NewCredential(secret)
	if _err == nil {
		return &openapi.Config{
			Credential: cred,
		}
	}
	return nil
}
func CreateImgClient() (_result *alimt20181012.Client, _err error) {
	configs := getBasicConfig()
	configs.Endpoint = tea.String("mt.cn-hangzhou.aliyuncs.com")
	_result = &alimt20181012.Client{}
	_result, _err = alimt20181012.NewClient(configs)
	return _result, _err
}

func CreateOCRClient() (_result *ocr_api20210707.Client, _err error) {
	config := getBasicConfig()
	config.Endpoint = tea.String("ocr-api.cn-hangzhou.aliyuncs.com")
	_result = &ocr_api20210707.Client{}
	_result, _err = ocr_api20210707.NewClient(config)
	return _result, _err
}

func TransImages(sourcePath string) error {
	dir, _err := tools.ListFilesInCurrentDirAny(sourcePath)
	if _err != nil {
		fmt.Printf("output: %v\n", _err)
		return _err
	}

	for _, s := range dir {
		TransImage(s)
	}

	return nil

}

func TransImage(filename string) (_err error) {

	name := path.Base(filename)
	dir := path.Dir(filename)
	client, _err := CreateImgClient()
	if _err != nil {
		return _err
	}

	file, err2 := os.ReadFile(filename)
	if err2 != nil {
		fmt.Printf("output: %v\n", err2)
		return
	}
	imgBase64 := base64.StdEncoding.EncodeToString(file)
	s, t := "zh", "zh-tw"

	newpath := fmt.Sprintf("%v/%v", dir, t)

	err := tools.CreateDirIfNotExist(newpath)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return err
	}

	translateImageRequest := &alimt20181012.TranslateImageRequest{
		ImageBase64:    &imgBase64,
		SourceLanguage: tea.String(s),
		TargetLanguage: tea.String(t),
		//Field:          tea.String("general"),
		Field: tea.String("e-commerce"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		res, _err := client.TranslateImageWithOptions(translateImageRequest, runtime)
		if _err != nil {
			return _err
		}

		newName := fmt.Sprintf("%v/%v", newpath, name)
		os.WriteFile(newName, DownImage(*res.Body.Data.FinalImageUrl), 0600)
		fmt.Printf("download: %v done\n", newName)

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

func TransTxt(txt, target string) (_res string, _err error) {
	client, _err := CreateImgClient()
	if _err != nil {
		return
	}

	translateGeneralRequest := &alimt20181012.TranslateGeneralRequest{
		FormatType:     tea.String("text"),
		SourceLanguage: tea.String("zh"),
		TargetLanguage: tea.String(target),
		SourceText:     tea.String(txt),
		Scene:          tea.String("general"),
		Context:        tea.String(""),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		res, _err := client.TranslateGeneralWithOptions(translateGeneralRequest, runtime)
		if _err != nil {
			return _err
		}

		_res = *res.Body.Data.Translated
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return
		}
	}
	return
}

func DownImage(url string) []byte {
	newRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return nil
	}
	cli := http.Client{}
	do, err := cli.Do(newRequest)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return nil
	}

	all, err := io.ReadAll(do.Body)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return nil
	}
	return all
}

func OcrImage(filename string) (_res string, _err error) {
	client, _err := CreateOCRClient()
	if _err != nil {
		return
	}

	file, err2 := os.ReadFile(filename)
	if err2 != nil {
		fmt.Printf("output: %v\n", err2)
		return
	}

	recognizeBasicRequest := &ocr_api20210707.RecognizeBasicRequest{
		Body: bytes.NewBuffer(file),
	}

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		res, _err := client.RecognizeBasicWithOptions(recognizeBasicRequest, runtime)
		if _err != nil {
			return _err
		}

		var D OcrData
		_e = json.Unmarshal([]byte(*res.Body.Data), &D)
		fmt.Printf("output: %v\n", D.Content)
		_res = D.Content

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return
		}
	}
	return
}

func _main(txt, target string) (_err error) {
	client, _err := CreateImgClient()
	if _err != nil {
		return _err
	}

	getBatchTranslateRequest := &alimt20181012.GetBatchTranslateRequest{
		FormatType:     tea.String("text"),
		SourceLanguage: tea.String("zh"),
		TargetLanguage: tea.String(target),
		Scene:          tea.String("general"),
		ApiType:        tea.String("translate_standard"),
		// { "11": "hello boy", "12": "go home", "13": "we can" }
		// 待翻译的条数不能超过 50
		// 单条翻译字符数不能超过 1000 字符
		// 总字符数不能超过 8000 字符
		// key 不会计入翻译的字符
		// 待翻译的内容中，标点、空格、html 标签均会计入字符
		SourceText: tea.String(txt),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.GetBatchTranslateWithOptions(getBatchTranslateRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
