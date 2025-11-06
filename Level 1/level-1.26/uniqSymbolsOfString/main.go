package main

import (
	"fmt"
	"strings"
)

// Примечание: bool - проще и понятнее для данного задания
// не стал юзать struct{}{}
// Оптимизация с struct{} в production, где важна каждая копейка памяти
func uniqSymbols(str string) bool {
	// для регистронезависимой проверки
	str = strings.ToLower(str)
	// map[rune]bool - ключ это символ, значение bool указывает на наличие
	uniqMap := make(map[rune]bool)
	// Перебираем каждый символ (руну) в строке
	for _, v := range str {
		// Если символ уже встречался, возвращаем false
		if uniqMap[v] {
			return false
		}
		// Отмечаем символ как встреченный(уже был)
		uniqMap[v] = true
	}
	// Если прошли всю строку и без повторений - все символы уникальны
	return true

}

func main() {
	examples := []string{"abcd", "abCdefAaf", "aabcd", "Wildberries", "Hello", "World"}

	for _, example := range examples {
		result := uniqSymbols(example)
		fmt.Printf("Строка: \"%s\" -> %t\n", example, result)
	}
}
