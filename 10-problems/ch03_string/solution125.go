package main

import (
	"fmt"
	"unicode"
)

func isalnum(ch rune) bool {
	return unicode.IsDigit(ch) || unicode.IsLetter(ch)
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	r := []rune(s)
	for i < j {
		for i < j && !isalnum(r[i]) {
			i++
		}
		for j > i && !isalnum(r[j]) {
			j--
		}
		if unicode.ToLower(r[i]) != unicode.ToLower(r[j]) {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func main() {
	// str1 := "A man, a plan, a canal: Panama"
	str1 := "0P"
	fmt.Println(isPalindrome((str1)))
}
