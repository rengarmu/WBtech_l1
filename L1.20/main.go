package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimRight(text, "\r\n") // удаляем \r и \n
	rs := reverseWords(text)
	fmt.Println(rs)
}

func reverseWords(text string) string {
	words := strings.Split(text, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
