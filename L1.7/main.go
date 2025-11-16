package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// NewSafeMap создает новый экземпляр SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int), // Инициализация пустой map
	}
}

// Set безопасно устанавливает значение по ключу
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()         // Блокировка для записи
	defer sm.mu.Unlock() // Гарантируем разблокировку при выходе из функции
	sm.data[key] = value //Запись значения в map
}

// Get безопасно получает значение по ключу
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key] // Получение значения по ключу
	return value, exists
}

// Delete безопасно удаляет значение по ключу
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key) // Удаление значения из map
}

// Len безопасно возвращает количество элементов в map
func (sm *SafeMap) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data) // Возвращаем количество элементов
}

// Keys безопасно возвращает все ключи из map
func (sm *SafeMap) Keys() []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// Создаем срез с начальной емкостью равной размеру map
	keys := make([]string, 0, len(sm.data))

	// Итерируем по всем ключам map
	for key := range sm.data {
		keys = append(keys, key) // Добавляем ключ в срез
	}
	return keys // Возвращаем срез ключей
}

// Clear безопасно очищает map
func (sm *SafeMap) Clear() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data = make(map[string]int) // Заменяем map на новую пустую
}

func main() {
	// Потокобезопасная map
	safeMap := NewSafeMap()

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутины для записи
	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go func(i int) {
			defer wg.Done() // Уменьшаем счетчик при завершении
			key := fmt.Sprintf("key%d", i)
			val := i
			safeMap.Set(key, val) // Записываем значение в map
			fmt.Printf("Set %s: %d\n", key, val)
		}(i) // Передаем i как параметр чтобы избежать замыкания на одной переменно
	}

	// Запускаем несколько горутин для чтения данных
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)

			// Безопасно читаем из map
			if value, exists := safeMap.Get(key); exists {
				fmt.Printf("Get %s: %d\n", key, value)
			} else {
				fmt.Printf("Key %s not found\n", key)
			}
		}(i)
	}

	// Запускаем несколько горутин для удаления данных
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			safeMap.Delete(key) // Безопасно удаляем из map
			fmt.Printf("Deleted %s\n", key)
		}(i)
	}
	// Ждём завершения всех горутин
	wg.Wait()

	// Выводим финальное состояние map
	fmt.Printf("\nFinal element count: %d\n", safeMap.Len())
	fmt.Printf("All Keys: %v\n", safeMap.Keys())

	//go run -race main.go
}
