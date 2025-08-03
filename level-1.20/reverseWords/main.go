package main

import "fmt"

func ReverseWordsCorrect(s string) string {
	if s == "" {
		return s
	}

	// Единственный дополнительный срез - копия исходной строки
	// так как строки immutable
	runes := []rune(s)

	// 1: Разворачиваем всю строку
	reverseRange(runes, 0, len(runes)-1)

	// 2: Разворачиваем каждое слово
	start := 0 // Начинаем с начала строки
	// Используем O(1) памяти, так как start - это просто индекс
	for i := 0; i <= len(runes); i++ { // Проходим по всем символам строки O(n)
		if i == len(runes) || runes[i] == ' ' { // Проверяем конец строки или пробел
			// Если нашли пробел или конец строки, разворачиваем слово
			// Если start < i, значит есть слово для разворота
			// Это условие гарантирует, что мы не разворачиваем пустые слова
			if i > start {
				reverseRange(runes, start, i-1)
			}
			start = i + 1 // Начинаем новое слово после пробела
		}
	}

	return string(runes) // Возвращаем результат как строку
}

// Вспомогательная функция - тоже O(1) памяти
func reverseRange(runes []rune, start, end int) {
	// Используем только локальные переменные O(1) размера
	for start < end { // start, end - O(1)
		// Обмен элементов без дополнительной памяти
		runes[start], runes[end] = runes[end], runes[start]
		start++ // O(1)
		end--   // O(1)
	}
}

func main() {
	// Тестируем правильный и неправильный подходы
	input := "snow dog sun"

	fmt.Printf("Исходная строка: %q\n", input)
	fmt.Printf("разворот: %q\n", ReverseWordsCorrect(input))

}
