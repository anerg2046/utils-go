package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Contain 检查是否包含
func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

// Pretty 友好显示控制台输出数据
func Pretty(data interface{}) {
	src, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(src))
}
