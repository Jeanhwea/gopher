package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

////////////////////////////////////////////////////////////////////////////////
// 获取错误的结构体，然后使用结构体中的参数进一步分析错误
func test01() {
	f, err := os.Open("test.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

////////////////////////////////////////////////////////////////////////////////
// 获取错误的结构体，然后使用结构中的方法进一步分析
func test02() {
	addr, err := net.LookupHost("notfound123.com")
	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)
}

////////////////////////////////////////////////////////////////////////////////
// 直接比较
func test03() {
	files, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic error:", err)
		return
	}
	fmt.Println("matched files", files)
}

func main() {
	// test01()
	// test02()
	test03()
}
