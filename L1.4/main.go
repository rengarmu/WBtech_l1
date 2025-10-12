package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/* Преимущества использования контекста (context):
- контекст является частью стандартного набора инструментов Go
- легко передать сигнал отключения сразу всем горутинам, независимо от уровня их вложения
- контекст поддерживает управление таймаутами
- горутины сами отвечают за проверку статуса контекста и прекращение своей активности, что минимизирует риск утечек ресурсов и зависающих потоков
*/

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done(): //Проверяем, не завершился ли контекст
			fmt.Println("Получен сигнал завершения. Завершаем работу...")
			return
		default:
			// Если сигнал о завершении не был получен
			fmt.Printf("Горутина %d выполняет работу...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	//WaitGroup для учета количества горутин
	var wg sync.WaitGroup

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	// Ловим сигнал прерывания (Ctrl + C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ждём прихода сигнала прерывания
	<-sigChan

	// Сигнал пришёл, отменяем контекст
	cancel()

	wg.Wait()

	fmt.Println("Все горутины успешно завершены.")
	os.Exit(0)
}
