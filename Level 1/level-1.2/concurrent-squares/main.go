package main

import "fmt"

func main() {

	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int)

	// Запускаем горутины
	for _, number := range numbers {
		go func(n int) {
			results <- n * n // Канал блокирует до получение
		}(number)
	}

	// Получаем результаты
	for range numbers {
		result := <-results // Канал блокирует до отправки
		fmt.Println(result)
	}
}
