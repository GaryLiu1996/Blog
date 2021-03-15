package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(GetAppPath())
	fmt.Println("hello world, guys!")
}
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

//E:\codes\私人\Blog\Program_Language\Golang\煎鱼\CH1\01_go的相对路径>go run main.go
//C:\Users\siyuan.liu\AppData\Local\Temp\go-build198123534\b001\exe
//hello world, guys!

//E:\codes\私人\Blog\Program_Language\Golang\煎鱼\CH1\01_go的相对路径>main.exe
//E:\codes\私人\Blog\Program_Language\Golang\煎鱼\CH1\01_go的相对路径
//hello world, guys!
