package tools

import (
	"fmt"
	"regexp"
	"strings"
)

/**
@description
@date: 04/04 17:30
@author Gk
**/

func GetErrHeader(hearder string) string {
	return "\"header\": {\"template\": \"orange\" ,   \"title\": {\"tag\": \"plain_text\",\"content\": \"" + hearder + "\"  }}"
}

func GetGreenHeader(hearder string) string {
	return "\"header\": {\"template\":\"green\",\"title\":{\"tag\":\"plain_text\",\"content\":\"" + hearder + "\"}}"
}

func GetHeader(hearder, style string) string {
	stylemap := map[string]string{
		"blue": "1", "wathet": "1", "turquoise": "1", "green": "1", "yellow": "1", "orange": "1", "red": "1", "carmine": "1", "violet": "1", "purple": "1", "indigo": "1", "grey": "1",
	}
	if stylemap[style] != "1" {
		style = "blue"
	}
	return "\"header\": {\"template\":\"" + style + "\",\"title\":{\"tag\":\"plain_text\",\"content\":\"" + hearder + "\"}}"
}

func GetErrContent(content ...string) string {
	tmp := fmt.Sprintf("{ \"tag\": \"div\", \"text\": { \"tag\": \"lark_md\", \"content\": \"**请检查输入或者确认是否有相关工单信息** \\n %s\"}}", strings.Join(content, ","))
	return fmt.Sprintf("\"elements\": [%v]", tmp)
}

func GetContent(subtitle string, content string) string {
	tmp := fmt.Sprintf("{ \"tag\": \"div\", \"text\": { \"tag\": \"lark_md\", \"content\": \"**%s** \\n %s\"}}", subtitle, content)
	return fmt.Sprintf("\"elements\": [%v]", tmp)
}

func GetCustomContent(content ...string) string {
	tmp := fmt.Sprintf("{ \"tag\": \"div\", \"text\": { \"tag\": \"lark_md\", \"content\": \"**%s** \"}}", strings.Join(content, ","))
	return fmt.Sprintf("\"elements\": [%v]", tmp)
}

func SplitLine() string {
	return "{\"tag\": \"hr\"}"
}

func GetMDByTag(key string, val interface{}) string {
	return fmt.Sprintf("{\"tag\": \"markdown\", \"content\": \"**%v:** %v\"}", key, val)
}

func GetMDByTagGen(val string) string {
	return fmt.Sprintf("{\"tag\": \"markdown\", \"content\": \"%v\"}", val)
}

// 页脚 | 备注
func GetNote(content string) string {
	return fmt.Sprintf("{\"tag\":\"note\",\"elements\":[%v]}", content)
}

// 页脚 | 备注, 放了一个 md的[]()
func GetNodeByMDUrl(name, url string) string {
	return fmt.Sprintf("{\"tag\":\"note\",\"elements\":[{\"tag\":\"lark_md\",\"content\":\"[%s](%v)\"}]}", name, url)
}

func GetFooter(content string) string {
	return fmt.Sprintf("{\"tag\":\"note\",\"elements\":[{\"tag\":\"lark_md\",\"content\":\"%v\"}]}", content)
}

// md 的[]() 不能直接使用
func GetContentMDUrl(name, url string) string {
	return fmt.Sprintf("[%s](%v)", name, url)
}

func GetDivMDByTag(content string) string {
	// 如果开头是4个空格会解析错误
	content = string(regexp.MustCompile("^(\\n)?(\\t)?(\\s){4}").ReplaceAll([]byte(content), []byte("&nbsp;\\t")))
	return fmt.Sprintf("{\"tag\": \"div\", \"text\": { \"tag\": \"lark_md\", \"content\": \"%v\"}}", content)
}

func GetCustomDivMDByTag(sep string, content ...string) string {
	return fmt.Sprintf("{\"tag\": \"div\", \"text\": { \"tag\": \"lark_md\", \"content\": \"%v\"}}", strings.Join(content, sep))
}

// 双列 是fields 嵌套GetDivFieldsByTag
func GetMoreCols(content ...string) string {
	return fmt.Sprintf("{\"tag\": \"div\",\"fields\":[%v]}", strings.Join(content, ","))
}

// isShort=true 占用一整行, false 占用半行
func GetDivFieldsByTag(content interface{}) string {
	return fmt.Sprintf("{\"is_short\":true,\"text\":{\"tag\":\"lark_md\",\"content\":\"%v\"}}", content)
}

func GetDivFieldsShortByTag(content interface{}, isShort bool) string {
	return fmt.Sprintf("{\"is_short\":%v,\"text\":{\"tag\":\"lark_md\",\"content\":\"%v\"}}", isShort, content)
}

func GetFmtTxt(header, val string) string {
	return fmt.Sprintf("{%s,%s}", header, val)
}

func GetFmtCard(header string, content ...string) string {
	return fmt.Sprintf("{%s,\"elements\" :[%s]}", header, strings.Join(content, ","))
}

// 消息卡片展示图片
// title: 图片标题 非必填
// content: 鼠标alt提示 非必填
func GetImage(title, content, key string) string {
	if title == "" {
		return fmt.Sprintf("{\"tag\":\"img\",\"img_key\":\"%s\",\"alt\":{\"tag\":\"plain_text\",\"content\":\"%s\"}}", key, content)
	}

	if content == "" {
		content = title
	}
	return fmt.Sprintf("{\"tag\":\"img\",\"title\":{\"tag\":\"lark_md\",\"content\":\"%s\"},\"img_key\":\"%s\",\"alt\":{\"tag\":\"plain_text\",\"content\":\"%s\"}}", title, key, content)
}

// GetFileFormat 获取一个格式化后的文件消息, 用于使用机器人发送文件消息
func GetFileFormat(fileKey string) string {
	return fmt.Sprintf("{\"file_key\":\"%s\"}", fileKey)
}

func GetVolcFormat(key string, val interface{}) string {
	return fmt.Sprintf("{\"is_short\":true,\"text\":{\"tag\":\"lark_md\",\"content\":\"**%v:** %v\"}}", key, val)
}
