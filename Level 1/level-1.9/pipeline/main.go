package main

import (
	"fmt"
	"sync"
)

// generateNumbers - генератор чисел (этап 1)
func generateNumbers(numbers []int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)

	for _, num := range numbers {
		fmt.Printf("Генератор - Отправляем: %d\n", num)
		output <- num
	}
	fmt.Println("Генератор -Завершен")

}

// processNumbers - обработчик чисел (этап 2)
func processNumbers(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)

	for num := range input {
		result := num * 2
		fmt.Printf("[Обработчик] %d * 2 = %d\n", num, result)
		output <- result
	}
	fmt.Println("[Обработчик] Завершен")
}

// outputResults - вывод результатов (этап 3)
func outputResults(input <-chan int) {
	fmt.Println("[Вывод] результатов")
	for result := range input {
		fmt.Printf("Результат: %d\n", result)
	}
	fmt.Println("[Вывод] Завершен")
}

// main - координатор конвейера
func main() {
	fmt.Println("Запуск конвейера чисел")
	// Исходные данные
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Входные данные: %v\n\n", numbers)

	// Создаем каналы
	firstChannel := make(chan int)
	secondChannel := make(chan int)

	// Группа ожидания для синхронизации
	var wg sync.WaitGroup

	// Запускаем этап 1: генерация
	wg.Add(1)
	go generateNumbers(numbers, firstChannel, &wg)

	// Запускаем этап 2: обработка
	wg.Add(1)
	go processNumbers(firstChannel, secondChannel, &wg)

	// Горутина для корректного закрытия каналов
	go func() {
		wg.Wait() // Ждем завершения этапов 1 и 2
	}()

	// Запускаем этап 3: вывод (в main goroutine)
	outputResults(secondChannel)

	fmt.Println("Конвейер завершен")
}
