package utils

import "log"

// CheckErr 错误检查
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
