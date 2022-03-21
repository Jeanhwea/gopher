package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) calcArea0(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("rect %v's length < 0\n", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width < 0\n", r)
		wg.Done()
		return
	}

	area := r.length * r.width
	wg.Done()
	fmt.Printf("rect %v's area = %d\n", r, area)
}

func (r rect) calcArea(wg *sync.WaitGroup) {
	// 避免每次 return 前都需要写 wg.Done(), 直接使用 defer 来确保锁释放
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length < 0\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width < 0\n", r)
		return
	}

	area := r.length * r.width
	fmt.Printf("rect %v's area = %d\n", r, area)
}

func main() {
	var wg sync.WaitGroup
	rects := []rect{
		{-67, 89},
		{7, -9},
		{7, 9},
	}

	for _, rect := range rects {
		wg.Add(1)
		go rect.calcArea(&wg)
	}
	wg.Wait()
	fmt.Println("Done all calculation!")
}
