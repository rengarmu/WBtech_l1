package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 1. Выход по условию
func exitByCondition(wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счётчика WaitGroup при завершении горутины
	i := 0
	for {
		if i >= 5 {
			fmt.Println("Exit by condition")
			return
		}
		fmt.Printf("Сondition: %d\n", i)
		i++
		time.Sleep(1 * time.Second) // Задержка в 1 секунду
	}
}

// 2. Через канал уведомления
func exitByChannel(stop <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // Всегда уменьшаем счётчик WaitGroup при завершении

	for {
		select {
		case <-stop:
			// Получен сигнал остановки через канал
			fmt.Println("Exit by channel")
			return
		default:
			// Если сигнал не получен, продолжаем работу
			fmt.Println("Channel: working...")
			time.Sleep(1 * time.Second) // Задержка в 1 секунду
		}
	}
}

// 3. Через контекст
func exitByContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счётчик WaitGroup при завершении
	for {
		select {
		case <-ctx.Done():
			// Получен сигнал остановки через контекст
			fmt.Printf("Exit by context: %v\n", ctx.Err())
			return
		default:
			// Если сигнал не получен, продолжаем работу
			fmt.Println("Context: working...")
			time.Sleep(1 * time.Second) // Задержка в 1 секунду
		}
	}
}

// 4. Использование runtime.Goexit()
func exitByGoexit(wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счётчика WaitGroup при завершении горутины

	//defer функции выполняются даже при вызове runtime.Goexit()
	defer fmt.Println("Defer executed before Goexit")

	// Запускаем внутреннюю горутину, которая вызовет Goexit()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine completed internally")
		// runtime.Goexit() завершает ТОЛЬКО текущую горутину, не всю программу
		runtime.Goexit()
	}()

	// Эта часть выполняется в основной горутине функции
	for i := 0; i < 3; i++ {
		fmt.Println("Goexit: working...")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Goexit: exiting")
}

// 5. Паника и восстановление
func exitByPanic(wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счётчика WaitGroup при завершении горутины

	// recover() перехватывает панику и позволяет gracefully завершить горутину
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recovered:", r)
		}
	}()

	counter := 0
	for {
		counter++
		fmt.Printf("Panic: %d\n", counter)
		time.Sleep(1 * time.Second)

		// Имитируем критическую ошибку, вызывающую панику
		if counter >= 3 {
			panic("Critical error occurred!") // Вызываем панику
		}
	}
}

// 6. Runtime.Goexit() в отдельной горутине
func demonstrateGoexit(wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счётчика WaitGroup при завершении горутины
	fmt.Println("Demonstrate Goexit...")

	innerWg := &sync.WaitGroup{}
	innerWg.Add(1)

	go func() {
		defer innerWg.Done()
		defer fmt.Println("Inner goroutine: defer before Goexit")

		time.Sleep(800 * time.Millisecond)
		fmt.Println("Inner goroutine: calling runtime.Goexit()")
		// Goexit завершает только эту внутреннюю горутину
		// Все defer функции будут выполнены перед завершением
		runtime.Goexit()
		// Этот код никогда не выполнится
		fmt.Println("This code will not execute")
	}()

	// Ждем завершения внутренней горутины
	innerWg.Wait()
	fmt.Println("Goexit demo completed")
}

func main() {
	fmt.Println("Main goroutine started")

	var wg sync.WaitGroup // WaitGroup для координации завершения горутин

	// 1. Выход по условию
	fmt.Println("\nExit by condition:")
	wg.Add(1) // Увеличиваем счетчик перед запуском горутины
	go exitByCondition(&wg)
	wg.Wait() // Ждем завершения горутины

	// 2. Через канал уведомления
	fmt.Println("\nExit by channel:")
	stop := make(chan bool)
	wg.Add(1)
	go exitByChannel(stop, &wg)
	time.Sleep(3 * time.Second) // Даем горутине поработать
	stop <- true                // Отправляем сигнал остановки через канал
	wg.Wait()

	// 3. Через контекст
	fmt.Println("\nExit by context:")
	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go exitByContext(ctx, &wg)
	time.Sleep(3 * time.Second) // Даем горутине поработать
	cancel()                    // Отменяем контекст - сигнал для завершения горутины
	wg.Wait()

	// 4. Использование runtime.Goexit()
	fmt.Println("\nExit by runtime.Goexit():")
	wg.Add(1)
	go exitByGoexit(&wg)
	wg.Wait()

	// 5. Паника и восстановление
	fmt.Println("\nExit by panic:")
	wg.Add(1)
	go exitByPanic(&wg)
	wg.Wait()

	// 6. Дополнительная демонстрация Goexit
	fmt.Println("\nDemonstrate Goexit:")
	wg.Add(1)
	go demonstrateGoexit(&wg)
	wg.Wait()

	// 7. Комбинированный пример с WaitGroup - управление группой горутин
	fmt.Println("\nCombined example with WaitGroup:")
	var comboWg sync.WaitGroup

	// Запускаем несколько горутин и ждем их завершения
	for i := 1; i <= 5; i++ {
		comboWg.Add(1)
		go func(id int) {
			defer comboWg.Done()
			for j := 0; j <= 3; j++ {
				fmt.Printf("Goroutine %d: %d\n", id, j)
				time.Sleep(1 * time.Second)
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}
	comboWg.Wait() // Ждем завершения всех горутин в группе
	fmt.Println("All goroutines completed")
}
