package main

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func doWork(str string) {
	log.Println(str)
	time.Sleep(1 * time.Second)
}

func main() {
	// 使用 go 关键字加方法调用来启动协程
	go doWork("Worker#1")
	go doWork("Worker#2")
	time.Sleep(1 * time.Second)
	go doWork("Worker#3")
	time.Sleep(3 * time.Second)
	log.Println("Done")
}
