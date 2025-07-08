package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

/**
@description
@date: 03/16 15:56
@author Gk
**/

// MemData global data
var MemData []string

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func ReplaceSpecStr(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}

// Get 请求参数中特殊符号转义
func ReplaceUrlStr(s string) string {
	// %先, 不然会改掉其他的结果
	s = strings.ReplaceAll(s, "%", "%25")
	s = strings.ReplaceAll(s, "+", "%2B")
	s = strings.ReplaceAll(s, " ", "%20")
	s = strings.ReplaceAll(s, "/", "%2F")
	s = strings.ReplaceAll(s, "?", "%3F")
	s = strings.ReplaceAll(s, "#", "%23")
	s = strings.ReplaceAll(s, "&", "%26")
	s = strings.ReplaceAll(s, "=", "%3D")
	return s
}

// If 模拟三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// 发版状态
func GetProcessingStatus(val string) string {
	res := val
	if val == "1" {
		res = "处理中"
	} else if val == "4" {
		res = "处理完成"
	} else if val == "5" {
		res = "发版中"
	} else if val == "9" {
		res = "发版完成"
	}
	return res
}

func DecryptStr(encrypt string, key string) (string, error) {
	buf, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", fmt.Errorf("base64StdEncode Error[%v]", err)
	}
	if len(buf) < aes.BlockSize {
		return "", errors.New("cipher  too short")
	}
	keyBs := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(keyBs[:sha256.Size])
	if err != nil {
		return "", fmt.Errorf("AESNewCipher Error[%v]", err)
	}
	iv := buf[:aes.BlockSize]
	buf = buf[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(buf)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(buf, buf)
	n := strings.Index(string(buf), "{")
	if n == -1 {
		n = 0
	}
	m := strings.LastIndex(string(buf), "}")
	if m == -1 {
		m = len(buf) - 1
	}
	return string(buf[n : m+1]), nil
}

// UniqueStrings 去重字符串数组
func UniqueStrings(strs []string) []string {
	result := make([]string, 0)
	mapStr := make(map[string]bool)

	for _, str := range strs {
		if !mapStr[str] {
			result = append(result, str)
			mapStr[str] = true
		}
	}
	return result
}
