package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name   string `info:"name"`
	Gender string `info:"gender"`
}

func findTag(obj interface{}) {
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagname := t.Field(i).Name
		taginfo := t.Field(i).Tag.Get("info")
		fmt.Println(tagname, taginfo)
	}
}

func main() {
	var res resume
	findTag(&res)
}
