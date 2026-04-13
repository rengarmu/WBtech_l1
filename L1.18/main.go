package main

import (
	"fmt"
	"sync"
)

// Counter структура счётчика с защитой мьютексом
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc увеличивает счётчик на 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value возвращает текущее значение счётчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	count := Counter{}

	goroutines := 100
	incPerGoroutine := 200

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incPerGoroutine; j++ {
				count.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Итого: %d\n", count.Value())
}
