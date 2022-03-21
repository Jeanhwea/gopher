package main

import (
	"fmt"
	"strings"
	"unicode"
)

func printBytes(str string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()
}

func printChars(str string) {
	fmt.Printf("Chars: ")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()
}

func charTest() {
	// 字符串函数
	fmt.Println(unicode.IsDigit('9'))
	fmt.Println(unicode.IsLower('a'))
	fmt.Println(unicode.IsLetter('d'))

	// 变量字符串
	name := "王大锤"
	for index, rune := range name {
		fmt.Printf("%d %c\n", index, rune)
	}
	fmt.Println(name[0])

	// byte 转字符串
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)

	// 字符串转 byte 数组
	str1 := "hello world"
	arr1 := []byte(str1)
	fmt.Printf("len of byte array = %d\n", len(arr1))
}

func stringTest() {
	str1 := "hello world"
	idx := strings.Index(str1, "w")
	fmt.Println(idx)

	var mutiline = `
		this is a mutiline string.
		this is second line
		and third line`
	fmt.Println(mutiline)
}

func stringsFunctions() {
	str1 := "apple banana"
	strings.Contains(str1, "ban")

	// startWith
	strings.HasPrefix("hello world", "prefix")
	// endWith
	strings.HasSuffix("hello world", "suffix")

	// replace
	str2 := "hello world"
	str3 := strings.Replace(str2, "world", "tom", 99)
	fmt.Printf("str2 = %v, str3 = %v", str2, str3)
}

func main() {
	// str1 := "hello world"
	// printBytes(str1)
	// printChars(str1)
	// charTest()
	// stringTest()
	stringsFunctions()
}
