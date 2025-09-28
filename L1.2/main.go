package main

import (
	"fmt"
	"sync"
)

func main() {
	// Исходный массив чисел
	nums := []int{2, 4, 6, 8, 10}
	// Создаем массив для записи результатов вычислений
	squares := make([]int, len(nums))

	// WaitGroup используется для ожидания завершения всех горутин
	var wg sync.WaitGroup

	for i, num := range nums {
		// Увеличиваем счетчик WaitGroup на 1 для каждой запускаемой горутины
		wg.Add(1)

		// Запускаем функцию как горутину
		go func(j, n int) {
			// Уменьшаем счетчик WaitGroup при завершении горутины
			defer wg.Done()
			// Вычисляем квадрат числа и сохраняем результат
			squares[j] = n * n
		}(i, num) // Передаем текущие значения i и num в горутину
	}
	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим результаты в исходном порядке
	for i, num := range nums {
		fmt.Printf("%d^2: %d\n", num, squares[i])
	}
}
