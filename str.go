package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// Int64ToStr 64整型转字符串
func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Float64ToStr 浮点转字符串
func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'g', 12, 64)
}

// StrToInt64 字符串转int64
func StrToInt64(s string) int64 {
	d, _ := strconv.ParseInt(s, 10, 64)
	return d
}

// StrToInt32 字符串转int32
func StrToInt32(s string) int32 {
	d, _ := strconv.ParseInt(s, 10, 32)
	return int32(d)
}

// StrToFloat64 字符串转浮点
func StrToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// InterfaceToString 接口类型转字符串
func InterfaceToString(data interface{}) string {
	var ret string
	if reflect.TypeOf(data).String() == "float64" {
		ret = Float64ToStr(data.(float64))
	} else if reflect.TypeOf(data).String() == "int64" {
		ret = Int64ToStr(data.(int64))
	} else if reflect.TypeOf(data).String() == "[]uint8" {
		ret = Bytes2String(data.([]uint8))
	} else if reflect.TypeOf(data).String() == "string" {
		ret = data.(string)
	} else {
		fmt.Println("未识别的类型：", reflect.TypeOf(data).String())
		ret = ""
	}
	return ret
}

// InterfaceToInt64 接口类型转int64
func InterfaceToInt64(data interface{}) int64 {
	var ret int64
	if reflect.TypeOf(data).String() == "float64" {
		ret = int64(data.(float64))
	} else if reflect.TypeOf(data).String() == "string" {
		ret = StrToInt64(data.(string))
	} else {
		ret = data.(int64)
	}
	return ret
}

// Bytes2String []byte转为string
func Bytes2String(bs []uint8) string {
	ba := make([]byte, 0, len(bs))
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
