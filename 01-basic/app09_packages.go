////////////////////////////////////////////////////////////////////////////////
// 常见的导包方式
// import (
//	"go/types"
//	"golang.org/x/syscall"
//	"errors"
//	xerrors "golang.org/x/errors"
//	_ "os/signal"
// )

package main

import "log"

// 包的 init 方法会在导入时调用
func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

// go mod init 初始化 go 的模块
// 包可以没有 main 方法
func main() {
	log.Println("Main runs!")
}
