package main

import "fmt"

// SetBit устанавливает i-й бит числа n в значение bit (0 или 1)
func SetBit(n int64, i uint, bit int) int64 {
	if bit == 1 {
		// Устанавливаем i-й бит в 1, операция ИЛИ
		return n | (1 << i)
	} else {
		// Устанавливаем i-й бит в 0, операция И с инвертированной маской
		return n &^ (1 << i)
	}
}

func PrintBinary(n int64) {
	for i := 63; i >= 0; i-- {
		if (n>>uint(i))&1 == 1 {
			fmt.Print("1") // Если бит установлен (равен 1)
		} else {
			fmt.Print("0") // Если бит не установлен (равен 0)
		}
		if i%8 == 0 && i > 0 {
			fmt.Print(" ") // Разделитель между байтами
		}
	}
	fmt.Println("₂")
}

func main() {
	var num int64 = 46       // Исходное число
	var bitPosition uint = 7 // Номер бита для изменения
	var newBit int = 1       // Значение, которое нужно установить (0 или 1)

	fmt.Printf("Исходное число: %d\n", num)
	PrintBinary(num)

	// Устанавливаем бит
	res := SetBit(num, bitPosition, newBit)

	fmt.Printf("\nУстановка %d-го бита в %d:\n", bitPosition, newBit)
	fmt.Printf("Результат: %d\n", res)
	PrintBinary(res)

}
