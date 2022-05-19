package main

import (
	"fmt"
	"log"
)

func init() {
	// 设置日志的格式
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	// 设置日志的前缀
	log.SetPrefix("Prefix ")
}

type Student struct {
	id   int
	name string
}

func logValue() {
	stu := &Student{
		id:   2832,
		name: "王新刚",
	}

	fmt.Printf("%%v = %v\n", stu)   // 只打印值
	fmt.Printf("%%+v = %+v\n", stu) // 打印键和值
	fmt.Printf("%%#v = %#v\n", stu) // 打印结构体名称，键和值

	// %v = &{2832 王新刚}
	// %+v = &{id:2832 name:王新刚}
	// %#v = &main.Student{id:2832, name:"王新刚"}
}

func main() {
	logValue()

	log.Println("Info")
	log.Fatal("Fatal message, program exit!")
	log.Println("Info 2")
}
