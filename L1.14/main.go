package main

import (
	"fmt"
	"reflect"
)

func main() {
	value := []interface{}{1, "hello", 3.23, true, nil,
		make(chan int),
		make(chan string),
		make(<-chan bool), // только для приёма
		make(chan<- int)}  // только для отправки

	for _, v := range value {
		fmt.Printf("Значение: %v, тип: %s\n", v, detectType(v))
	}
}

func detectType(val interface{}) string {
	// Обработка nil
	if val == nil {
		return "nil"
	}

	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		// Для всех остальных проверяем, является ли тип каналом
		typ := reflect.TypeOf(val)
		if typ.Kind() == reflect.Chan {
			return "chan"
		}
		return fmt.Sprintf("unknown type (%s)", typ.Kind())
	}
}
