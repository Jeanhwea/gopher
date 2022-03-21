package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId   int
	custId  int
	ordCode string
	enabled bool
}

func createQuery(obj interface{}) {
	t := reflect.TypeOf(obj)  // => 获取 obj 的类型 main.order
	v := reflect.ValueOf(obj) // => 获取 obj 的值列表
	k := t.Kind()             // => 获取种类, 例如: Struct, Unit8, String, Array 等
	fmt.Printf("Type[%s], Value[%s], Kind[%s]\n", t, v, k)
	// Type[main.order], Value[{%!s(int=123) %!s(int=222) H021 %!s(bool=false)}], Kind[struct]

	// 获取每个域
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("Field[%d]: Type[%T], Value[%v]\n", i, f, f)
		switch f.Kind() {
		case reflect.Int:
			fmt.Println("Integer Value:", f.Int())
		case reflect.String:
			fmt.Println("String Value:", f.String())
		default:
			fmt.Println("Unknown Value")
		}
	}
}

func main() {
	order01 := order{
		ordId:   123,
		custId:  222,
		ordCode: "H021",
	}
	createQuery(order01)
}
