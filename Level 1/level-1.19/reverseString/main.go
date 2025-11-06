package main

import "fmt"

// Функция разворота строки через срез рун (разворот рун сохраняет целостность символов Unicode)
func ReverseString(s string) string {
	// Преобразуем строку в срез рун для корректной работы с Unicode
	runes := []rune(s)

	// Разворачиваем срез рун "на месте"
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем обратно в строку
	return string(runes)
}

func main() {
	reversed := ReverseString("Hello, World!")
	fmt.Println(reversed)
}
