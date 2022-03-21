package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 一次读取文件到内存
	data, err := ioutil.ReadFile("/etc/hosts")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Printf("Content of file:\n%s", string(data))
}
