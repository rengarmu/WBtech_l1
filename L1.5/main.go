package main

import (
	"fmt"
	"time"
)

func main() {
	//Время работы программы в секундах
	N := 4
	dataChan := make(chan int)

	//Горутина для отправки данных в канал
	go func() {
		for i := 1; ; i++ {
			dataChan <- i
			//Таймер задержки между отправками
			time.Sleep(500 * time.Millisecond)
		}
	}()
	timeout := time.After(time.Duration(N) * time.Second)

	//Цикл чтения из канала
	for {
		// Конструкция select позволяет ждать на нескольких каналах одновременно
		// Она блокируется до тех пор, пока один из case не станет доступным
		select {
		// case срабатывает, когда в канале dataChan появляется новое значение
		case value := <-dataChan:
			fmt.Printf("Получено значение %d\n", value)
		case <-timeout:
			fmt.Println("Программа завершена!")
			return
		}
	}
}
