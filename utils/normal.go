package utils

import (
	"fmt"
	"os"
	"reflect"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IPCheck 判断是否为IPv4路径
func IPCheck(ip string) bool {
	regex := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regex, strings.Trim(ip, " ")); match {
		return true
	}
	return false
}

// ProcessBar 进度条
func ProcessBar(now int, all int) {
	str := "[Check] [" + bar(int((float64(now)/float64(all))*20), 20) + "] " + strconv.Itoa(now) + "/" + strconv.Itoa(all)
	fmt.Printf("\r%s", str)
	if now == all {
		fmt.Println()
	}
}

func bar(count, size int) string {
	str := ""
	for i := 0; i < size; i++ {
		if i < count {
			str += "="
		} else {
			str += " "
		}
	}
	return str
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 动态调用
func Call(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
    if len(params) != f.Type().NumIn() {
        return nil, errors.New("the number of input params not match!")
    }
    in := make([]reflect.Value, len(params))
    for k, v := range params {
        in[k] = reflect.ValueOf(v)
    }
    return f.Call(in), nil
}