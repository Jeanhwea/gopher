package main

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

func pointerAttr() {
	var x struct {
		a int64
		b bool
		c string
	}

	const M, N = unsafe.Sizeof(x.c), unsafe.Sizeof(x)
	fmt.Println(M, N) // 16 32

	fmt.Println(unsafe.Alignof(x.a)) // 8
	fmt.Println(unsafe.Alignof(x.b)) // 1
	fmt.Println(unsafe.Alignof(x.c)) // 8

	fmt.Println(unsafe.Offsetof(x.a)) // 0
	fmt.Println(unsafe.Offsetof(x.b)) // 8
	fmt.Println(unsafe.Offsetof(x.c)) // 16
}

// unsafe.Pointer 非类型的安全组指针的 6 种使用场景

func uc1Pointer() {
	var num1 float64 = 3.4
	// 修改值的类型 math.Float64bits(num1)
	var bits uint64

	// 将 float64 -> uint64, 但是不新分配底层空间
	bits = *(*uint64)(unsafe.Pointer(&num1))
	fmt.Println(bits) // => 4614838538166547251
}

func uc2Pointer() {
	type T struct{ a int }
	var t T
	fmt.Printf("%p\n", &t)
	fmt.Println(&t)
	// 这里将字符转成 uintptr 来打印日志输出
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t)))
}

func uc3Pointer() {
	type T struct {
		x bool
		y [3]int16
	}

	const N = unsafe.Offsetof(T{}.y)
	const M = unsafe.Sizeof(T{}.y[0])

	t := T{y: [3]int16{1, 2, 3}}
	p := unsafe.Pointer(&t)

	// 这里转化成 uintptr 进行算术运算
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M))
	fmt.Println(*ty2)

	addr := uintptr(p) + N + M
	// 省略一大段逻辑代码
	fmt.Println(*(*int16)(unsafe.Pointer(addr))) // 这里 addr 指的内存块可能会被 GC
}

func uc4Pointer() {
	fd, err := syscall.Open("/tmp/file.txt", syscall.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}

	var buf = make([]byte, 128)
	size := 8

	// Linux 底层的调用需要提供地址，这是可以使用 uintptr 来进行调用
	syscall.Syscall(syscall.SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(&buf)), uintptr(size))

}

func uc5Pointer() {
	// 用于 reflect.ValueOf 的值立即转值
	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
	fmt.Println(p)

	u := reflect.ValueOf(new(int)).Pointer()
	// 在这个时刻，处于存储在 u 中的地址处的内存块
	// 可能会被回收掉。
	p2 := (*int)(unsafe.Pointer(u))
	fmt.Println(p2)
}

func uc6Pointer() {
	hello := [5]byte{'H', 'e', 'l', 'l', 'o'}
	world := []byte("world")

	// 将一个 reflect.SliceHeader 或者 reflect.StringHeader 值的 Data 字段转换为
	// 非类型安全指针，以及其逆转换
	header := (*reflect.SliceHeader)(unsafe.Pointer(&world))
	header.Data = uintptr(unsafe.Pointer(&hello))

	header.Len = 2
	header.Cap = len(hello)
	fmt.Printf("%s\n", world) // He
	world = world[:cap(world)]
	fmt.Printf("%s\n", world) // Hello
}

func main() {
	// uc1Pointer()
	// uc2Pointer()
	// uc3Pointer()
	uc6Pointer()
}
