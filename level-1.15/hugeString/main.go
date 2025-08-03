package main

import (
	"fmt"
	"strings"
)

// justString будет содержать первые 100 символов огромной строки
var justString string

// createHugeString создает огромную строку заданного размера, заполненную символами 'a'
func createHugeString(size int) string {
	var sb strings.Builder
	sb.Grow(size) // Предварительно выделяем память для эффективности

	for i := 0; i < size; i++ {
		// Добавляем символ 'a' в строку
		sb.WriteByte('a')
	}

	return sb.String()
}

/*
someFunc демонстрирует правильное решение проблемы утечки памяти
при работе со срезами строк
*/
func someFunc() {
	// Создаем огромную строку размером 1024 байта
	v := createHugeString(1024)

	// создаем независимую копию первых 100 символов от исходной большой строки
	var builder strings.Builder
	// Сохраняем только первые 100 символов огромной строки в переменную justString
	builder.WriteString(v[:100])
	justString = builder.String()
	// Теперь большая строка v удалится сборщиком мусора  v = ""
}

func main() {
	someFunc()
	fmt.Println(len(justString))
}

/*
ПРОБЛЕМА исходного кода: justString = v[:100]

Негативные последствия:
1. Утечка памяти - вся строка (1024 байта) остается в памяти
2. Неэффективное использование памяти (в 10+ раз > нужного)
3. Потенциальное повреждение Unicode-символов при срезе по байтам
4. Накопление утечек при повторных вызовах

РЕШЕНИЕ: создание независимой копии через strings.Builder
*/
