package util

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

type Arr []map[string]string

// 获取sessionid
func GetSessionId() string {
	return fmt.Sprintf("%d%d", rand.Intn(1000), time.Now().Nanosecond())
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 生成随机字符串
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// 十六进制转换
func ToHex16(num int) string {
	if num == 0 {
		return "0"
	}
	s := "0123456789abcdef"

	ret := ""
	for num != 0 {
		index := num & 15
		ret = s[index:index+1] + ret
		num >>= 4
	}

	return ret
}

func JSONToMap(str string) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}
	return tempMap
}

func StringToArr(arr string) []map[string]interface{} {
	var a []map[string]interface{}
	// 将字符串反解析为数组
	json.Unmarshal([]byte(arr), &a)
	fmt.Println(a) // [1 2 3 4]
	return a
}

// 生成指定位数的uuid
func GetUuid(count int) string {
	u := uuid.NewV4().String()
	newUuid := strings.Replace(u, "-", "", -1)
	content := newUuid[0:count]
	return content
}

// 生成随机数
func GetRand() int {
	return int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

