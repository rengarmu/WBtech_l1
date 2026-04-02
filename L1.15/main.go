package main

import "strings"

/* Срез в Go ссылается на ту же область памяти, что и исходная строка,
поэтому justString удерживает в памяти всю исходную строку,
а не только первые 100 символов.
После завершения функции someFunc память, выделенная под v не освобождается,
так как на неё остаётся ссылка через justString.
Это приводит к утечке памяти.
*/

var justString string

func createHugeString(size int) string {
	return strings.Repeat("s", size)
}

func someFunc() {
	v := createHugeString(1 << 10)
	// Копируем только нужные 100 байт
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
