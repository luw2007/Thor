package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	ErrorCode      = -1
	ErrorIPBanCOde = -2
)

// RandStringBytesMaskImpr 随机生成长度为n的字符串
func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// Md5DigestString md5.digest
func Md5DigestString(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateSign 生成Sign验签字段
func GenerateSign(params map[string]string) string {
	sortKeys := make([]string, 0)
	for k := range params {
		sortKeys = append(sortKeys, k)
	}
	sort.Strings(sortKeys)
	sign := make([]string, len(sortKeys))
	for i, k := range sortKeys {
		sign[i] = k + params[k]
	}
	return strings.Join(sign, "")
}

func GetCode(code interface{}) int {
	switch ok := code.(type) {
	default:
		fmt.Println("err code type:", code)
		goto BadCode
	case float64:
		return int(ok)
	case int:
		return ok
	case string:
		c, err := strconv.Atoi(ok)
		if err != nil {
			fmt.Println("err code string:", ok)
			goto BadCode
		}
		return c
	}
BadCode:
	return ErrorCode
}

func IsStatusOk(code interface{}) bool {
	return GetCode(code) == 0
}
