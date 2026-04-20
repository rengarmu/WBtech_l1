package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Задаём значения a и b, каждое больше 2^20 = 1,048,576
	// Используем big.Int, чтобы избежать переполнения даже при очень больших числах
	a := new(big.Int)
	b := new(big.Int)

	// Устанавливаем a и b
	a.SetString("1048577", 10)
	b.SetString("1048578", 10)

	// Выводим исходные значения
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println()

	//Вычисления
	bigProd(a, b)
	bigQuot(a, b)
	bigSum(a, b)
	bigDiff(a, b)

}

// Сложение
func bigSum(a, b *big.Int) {
	sum := new(big.Int).Add(a, b)
	fmt.Printf("a + b = %v\n", sum)
}

// Вычитание
func bigDiff(a, b *big.Int) {
	diff := new(big.Int).Sub(a, b)
	fmt.Printf("a - b = %v\n", diff)
}

// Умножение
func bigProd(a, b *big.Int) {
	prod := new(big.Int).Mul(a, b)
	fmt.Printf("a * b = %v\n", prod)
}

// Деление
func bigQuot(a, b *big.Int) {
	if b.Cmp(big.NewInt(0)) != 0 {
		quot := new(big.Int).Div(a, b)
		fmt.Printf("a / b = %v\n", quot)
	} else {
		fmt.Println("Ошибка: деление на ноль")
	}
}
