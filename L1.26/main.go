package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	// Приводим строку к нижнему регистру
	lowStr := strings.ToLower(s)

	// Map для отслеживания встреченных символов
	seen := make(map[rune]bool)

	// Перебираем руны (символы Unicode)
	for _, char := range lowStr {
		if seen[char] == true {
			return false
		}
		seen[char] = true
	}
	return true
}

func main() {
	testCases := []string{
		"abcd",
		"abCdefAaf",
		"Hello",
		"Привет",
		"",
		"123 456",
		"GoLang",
		"😊👍😊",
	}
	for _, str := range testCases {
		fmt.Printf("%q -> %v\n", str, isUnique(str))
	}
}
