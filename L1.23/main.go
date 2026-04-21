package main

import "fmt"

// Remove удаляет элемент с индексом i из слайса s.
func Remove[T any](s []T, i int) []T {
	// Проверка границ (защита от паники)
	if i < 0 || i >= len(s) {
		return s
	}
	// Сдвигаем хвост [i+1:] на место i
	copy(s[i:], s[i+1:])
	// Обнуляем последний элемент (освобождаем ссылку для GC)
	var n T
	s[len(s)-1] = n
	// Укорачиваем слайс
	return s[:len(s)-1]
}

func main() {
	// Пример с int (значения)
	s := []int{1, 2, 3, 4, 5}
	s = Remove(s, 2)
	fmt.Println(s) // [1 2 4 5]

	// Пример с указателями (структура)
	type Words struct{ word string }
	s2 := []*Words{{"time"}, {"space"}, {"power"}}
	s2 = Remove(s2, 1)
	fmt.Println(s2[1]) // &{power}
}
