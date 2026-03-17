package main

import "fmt"

func main() {
	// Исходные множества, заданные в виде слайсов
	A := []int{1, 1, 2, 2, 3, 7}
	B := []int{2, 2, 3, 3, 4}

	// map[int]struct{} — пустая структура экономит память.
	set := make(map[int]struct{}, len(A))

	// Заполняем множество элементами из A, дубликаты игнорируются
	for _, v := range A {
		set[v] = struct{}{}
	}

	// Результирующий слайс для пересечения
	res := make([]int, 0, min(len(A), len(B)))
	// Проходим по второму слайсу B
	for _, i := range B {
		if _, ok := set[i]; ok {
			res = append(res, i)
			delete(set, i)

		}
	}
	fmt.Println(res)
}

// min — вспомогательная функция для нахождения минимума двух целых чисел.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
