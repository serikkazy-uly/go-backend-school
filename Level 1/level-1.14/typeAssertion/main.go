package main

import (
	"fmt"
	"reflect"
)

// https://go.dev/ref/spec#Type_switches

func main() {
	// Примеры различных типов
	variables := []interface{}{
		42,              // int
		"Hello, World!", // string
		true,            // bool
		make(chan int),  // chan
		3.14,            // float64
	}
	// Применяем typeAssertion к каждому элементу
	for _, v := range variables {
		typeAssertion(v)
	}
}

// typeAssertion принимает переменную типа interface{} и определяет её тип
func typeAssertion(i interface{}) {
	// Используем type switch для определения типа переменной
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	case chan int:
		fmt.Printf("Channel: (type: %T)\n", v)

	default:
		fmt.Printf("Type: %v\n", reflect.TypeOf(i)) // Используем reflect для получения типа
	}
}
