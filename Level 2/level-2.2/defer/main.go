package main

import "fmt"

// x - именованная возвращаемая переменная
// x - это переменная в области видимости функции
// x - это также возвращаемое значение
// defer может изменить x, и это повлияет на результат
func test() (x int) {
	defer func() {
		x++ // изменяет возвращаемую переменную x
	}()
	x = 1

	return // возвращает значение переменной x
}

// НЕименованное возвращаемое значение "Не могут" быть изменены в defer
// defer может изменить локальную x, но не возвращаемое значение
func anotherTest() int {
	var x int // локальная переменная
	// fmt.Println("anotherTest() - var x : ", x) // 0
	defer func() {
		x++ // изменяет локальную переменную x
		// fmt.Println("anotherTest() - inside defer: ", x) // 2
	}()
	x = 1
	// fmt.Println("anotherTest() - after defer: ", x) // 1
	// return x копирует значение x в возвращаемое значение
	return x //  возвращает КОПИЮ значения x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

/*
Выввод:
2 - test() // именованная возвращаемая переменная
1 - anotherTest() // НЕименованное возвращаемое значение "Не могут" быть изменены в defer

*/
