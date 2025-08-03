package main

import "fmt"

// Итеративная реализация бинарного поиска
func binarySearchIterative(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		// Вычисляем средний индекс (избегаем переполнения)
		mid := left + (right-left)/2

		// Если нашли элемент
		if arr[mid] == target {
			return mid
		}

		// Если искомый элемент меньше среднего - ищем в левой половине
		if target < arr[mid] {
			right = mid - 1
		} else {
			// Иначе ищем в правой половине
			left = mid + 1
		}
	}

	// Элемент не найден
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19} // Отсортированный массив
	targets := []int{7, 10, 1, 19, 20, 0}           // Элементы для поиска

	fmt.Println("Отсортированный массив:", arr)
	fmt.Println()

	for _, target := range targets {
		result := binarySearchIterative(arr, target)
		if result != -1 {
			fmt.Printf("Элемент %d найден по индексу %d\n", target, result)
		} else {
			fmt.Printf("Элемент %d не найден\n", target)
		}
	}
}
