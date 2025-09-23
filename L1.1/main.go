package main

import "fmt"

// Структура Human
type Human struct {
	Name string
	Age  int
}

// Структура Action имеет вложенную структуру Human и дополнительные поля
type Action struct {
	Human //Вложенная структура (embedded struct)
	Name  string
	Class int
}

func main() {
	// Создаем объект типа Action, передавая необходимые параметры
	act := buildingAction("Анастасия Иванова", 30, "бег", 3)
	// Выводим подробное представление структуры act
	fmt.Printf("%+v\n", act)

	// Получаем доступ к полю Age
	fmt.Println(act.Human.Age)
	fmt.Println(act.Age)

	// Вывод имени
	fmt.Println(act.Name)
	fmt.Println(act.Human.Name)

	// Вызываем метод printName() из встроенной структуры Human
	act.printName()

}

// Функция для создания экземпляра Action
func buildingAction(humanName string, age int, actionName string, class int) Action {
	return Action{
		Human: Human{humanName, age}, // Заполняем встроенное поле Human
		Name:  actionName,
		Class: class,
	}
}

// Метод для вывода имени из структуры Human
func (h Human) printName() {
	fmt.Println(h.Name)
}
