package main

import (
	"fmt"
	"runtime"
)
import "github.com/zhangyiming748/GetFileInfo"

func main() {
	var root string
	if runtime.GOOS == "android" && runtime.GOARCH == "arm64" {
		root = "sdcard/Movies/bili"
	}
	files := GetFileInfo.GetAllFilesInfo(root, "xml")
	for _, file := range files {
		fmt.Println(file)
	}
}
