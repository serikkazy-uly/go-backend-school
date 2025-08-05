package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

func RemovePersonElement(slice []*Person, idx int) []*Person {
	// индекс вне границ
	if idx < 0 || idx >= len(slice) {
		return slice
	}

	// Сдвигаем хвост слайса на место удаляемого элемента
	//  <- согласно idx (заданный элемент)
	copy(slice[idx:], slice[idx+1:])

	//  Обнуляем последний указатель для предотвращения утечки памяти
	slice[len(slice)-1] = nil // при иcпользовании int вмместо nil 0

	// Возвращаем слайс с уменьшенной длиной на 1
	return slice[:len(slice)-1]
}

func main() {
	fmt.Println("Удаление из слайса указателей:")
	people := []*Person{
		{"Алиса", 25},
		{"Дамир", 30},
		{"Ваня", 35},
		{"Диана", 40},
	}
	fmt.Printf("   Исходный: %v\n", people)
	people = RemovePersonElement(people, 1) // удаляем Дамира
	fmt.Printf("   Результат: %v\n", people)

}
