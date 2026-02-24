package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Создаем 2 канала для передачи данных

	input := make(chan int)
	output := make(chan int)

	// Первая горутина, куда пишутся числа x из массива
	go func() {
		for _, x := range nums {
			input <- x
		}
		close(input) // Закрываем канал после отправки всех чисел
	}()

	// Вторая горутина, куда пишется  результат операции x*2
	go func() {
		for x := range input {
			output <- x * 2
		}
		close(output) // Закрываем канал после отправки всех результатов
	}()

	// Выводим результаты
	for res := range output {
		fmt.Println(res)
	}
}
