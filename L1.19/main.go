package main

/*
Разработать программу, которая
переворачивает подаваемую на вход строку.

Например: при вводе строки «главрыба»
вывод должен быть «абырвалг».

Учтите, что символы могут быть в Unicode
(русские буквы, emoji и пр.), то есть просто
iterating по байтам может не подойти —
нужен срез рун ([]rune).

*/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimRight(text, "\r\n") // удаляем \r и \n
	rs := reverseStr(text)
	fmt.Println(rs)
}

func reverseStr(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
