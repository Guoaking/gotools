package ali

import (
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
func CreateClient() (_result *alimt20181012.Client, _err error) {
	// 工程代码建议使用更安全的无AK方式，凭据配置方式请参见：https://help.aliyun.com/document_detail/378661.html。

	config := new(credential.Config).
		SetType("access_key").
		// 从环境变量中获取AccessKey Id。
		SetAccessKeyId(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")).
		// 从环境变量中获取AccessKey Secret
		SetAccessKeySecret(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"))
	// 从环境变量中获取STS临时凭证。
	//SetSecurityToken(os.Getenv("ALIBABA_CLOUD_SECURITY_TOKEN"))

	cred, _err := credential.NewCredential(config)
	if _err != nil {
		return _result, _err
	}

	configs := &openapi.Config{
		Credential: cred,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/alimt
	configs.Endpoint = tea.String("mt.cn-hangzhou.aliyuncs.com")
	_result = &alimt20181012.Client{}
	_result, _err = alimt20181012.NewClient(configs)
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
	client, _err := CreateClient()
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

//curl 'https://cdn.translate.alibaba.com/r/1e55a01512a24f85b67c2317cff59fae.jpg' \
//  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' \
//  -H 'Accept-Language: zh-CN,zh;q=0.9' \
//  -H 'Cache-Control: no-cache' \
//  -H 'Connection: keep-alive' \
//  -b '_samesite_flag_=true; cookie2=19db3951ef496dbce32d92fb92bc4c41; t=096ee7eee9a93a67922f19e9e9faf27b; _tb_token_=e471e6e554e7e; ali_apache_id=33.50.128.217.1739786695497.720779.4; isg=BOXl1IbT7ldZlAqLXrPSjzam9KcfIpm0L344rOfKiZwr_gRwr3CshYeUiGKIfrFs' \
//  -H 'DNT: 1' \
//  -H 'Pragma: no-cache' \
//  -H 'Sec-Fetch-Dest: document' \
//  -H 'Sec-Fetch-Mode: navigate' \
//  -H 'Sec-Fetch-Site: none' \
//  -H 'Sec-Fetch-User: ?1' \
//  -H 'Upgrade-Insecure-Requests: 1' \
//  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36' \
//  -H 'sec-ch-ua: "Not)A;Brand";v="8", "Chromium";v="138"' \
//  -H 'sec-ch-ua-mobile: ?0' \
//  -H 'sec-ch-ua-platform: "macOS"' \
//  -H 'x-no-edit-page: 1'
