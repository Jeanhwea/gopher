package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func createTempDir() {
	dir, _ := ioutil.TempDir("", "taste-dir-*")
	fmt.Println(dir)

	resp, _ := os.MkdirTemp("", "taste-aaa")
	fmt.Println(resp)
}

func createTempFile() {
	file, _ := ioutil.TempFile("", "taste-file-*")
	fmt.Println(file.Name())

	fileinfos, _ := ioutil.ReadDir("~/pums_tmp")
	fmt.Printf("fileinfos = %+v", fileinfos)
	for _, fileinfo := range fileinfos {
		fmt.Println(fileinfo.Name())
	}
}

func main() {
	createTempFile()
	// createTempDir()
}
