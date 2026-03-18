package main

import (
	"fmt"
)

func main() {
	// Исходная последовательность строк
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаём множество с помощью map[string]struct{}
	set := make(map[string]struct{})

	// Добавляем каждое слово в множество: ключ — слово, значение — пустая структура
	for _, word := range words {
		set[word] = struct{}{}
	}

	//Собраем уникальные элементы в слайс
	unique := make([]string, 0, len(set))
	for word := range set {
		unique = append(unique, word)
	}

	fmt.Println(unique) // Вывод: [cat dog tree]
}
