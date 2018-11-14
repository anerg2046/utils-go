package helper

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetCurrentDirectory 获取当前路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
