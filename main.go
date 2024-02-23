package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"strings"
)
import "github.com/zhangyiming748/GetFileInfo"

func main() {
	var root string
	if runtime.GOOS == "android" && runtime.GOARCH == "arm64" {
		root = "/sdcard/Movies/bili"
	}
	files := GetFileInfo.GetAllFilesInfo(root, "xml")
	for _, file := range files {
		fmt.Println("找到的文件", file)
		ass := strings.Replace(file.FullPath, ".xml", ".ass", 1)
		//python /data/data/com.termux/files/home/bin/share/danmaku2ass/danmaku2ass.py "/sdcard/Movies/bili/报恩榴莲了解一下好想吃啊 这些榴莲是来报恩的吗为什么果肉那么多好想吃啊.xml" -s 1280x720 -dm 15 -o "/sdcard/Movies/bili/报恩榴莲了解一下好想吃啊 这些榴莲是来报恩的吗为什么果肉那么多好想吃啊.ass"
		xml2ass := exec.Command("python /data/data/com.termux/files/home/bin/share/danmaku2ass/danmaku2ass.py", file.FullPath, "-s", "1280x720", "-dm", "15", "-o", ass)
		output, err := xml2ass.CombinedOutput()
		if err != nil {
			slog.Warn("命令运行失败", slog.String("命令原文", fmt.Sprint(xml2ass)), slog.String("错误原文", fmt.Sprint(err)))
		} else {
			slog.Warn("命令运行成功", slog.String("命令原文", fmt.Sprint(xml2ass)), slog.String("输出原文", fmt.Sprint(string(output))))
			os.Remove(file.FullPath)
		}
	}
}
