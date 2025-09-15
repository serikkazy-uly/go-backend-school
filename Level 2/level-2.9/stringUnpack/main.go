package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// UnpackString выполняет распаковку строки с повторяющимися символами
func UnpackString(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	runes := []rune(s)
	var result strings.Builder

	// Проверяю корректность строки - не должна начинаться с цифы
	if unicode.IsDigit(runes[0]) {
		return "", errors.New("некорректная строка: начинается с цифры")
	}

	i := 0
	for i < len(runes) {
		char := runes[i]

		// Обработка escape последовательностей
		if char == '\\' && i+1 < len(runes) {
			i++ // пропускаем обратный слеш
			escapedChar := runes[i]

			// Проверяю есть ли после экранированного символа - цифра
			if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				digit := runes[i+1]
				count, _ := strconv.Atoi(string(digit))

				// Добавил экранированный символ count раз
				for j := 0; j < count; j++ {
					result.WriteRune(escapedChar)
				}
				i += 2 // пропускаю экранированный символ и цифру
			} else {
				// Добавляю экранированный символ один раз
				result.WriteRune(escapedChar)
				i++
			}
			continue
		}

		// Если текущий символ - цифра, это ошибка (должен быть после символа)
		if unicode.IsDigit(char) {
			// Проверяем случай подряд идущих цифр
			return "", errors.New("некорректная строка: некорректное использование цифр")
		}

		// Обработка обычных символов
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			// После символа идет цифра - распаковка
			digit := runes[i+1]
			count, _ := strconv.Atoi(string(digit))

			// Добавляю символ count раз
			for j := 0; j < count; j++ {
				result.WriteRune(char)
			}
			i += 2 // пропускаю символ и цифру
		} else {
			// Обычный символ без цифры
			result.WriteRune(char)
			i++
		}
	}

	return result.String(), nil
}

func main() {
	testCases := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
	}

	fmt.Println("Примеры работы функции UnpackString:")
	for _, test := range testCases {
		result, err := UnpackString(test)
		if err != nil {
			fmt.Printf("Вход: %q -> Ошибка: %v\n", test, err)
		} else {
			fmt.Printf("Вход: %q -> Выход: %q\n", test, result)
		}
	}
}
