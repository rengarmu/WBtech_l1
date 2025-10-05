package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Проверяем аргументы командной строки и парсим количество воркеров
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: не указано количество воркеров")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Ошибка: количество воркеров должно быть положительным числом")
		return
	}

	// Создаем канал для данных
	dataChannel := make(chan string)
	var wg sync.WaitGroup

	// Запускаем воркеры
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg)
	}

	// Главная горутина - постоянная запись данных в канал
	counter := 1
	for {
		message := fmt.Sprintf("#%d", counter)
		dataChannel <- message
		counter++
		time.Sleep(500 * time.Millisecond)
	}
}

// worker читает данные из канала и выводит в stdout
func worker(id int, dataChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range dataChannel {
		fmt.Printf("Worker %d: %s\n", id, data)
	}
}
