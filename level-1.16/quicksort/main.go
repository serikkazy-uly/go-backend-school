package main

import "fmt"

// Лучший/Средний случай O(n log n)
// Худший случай O(n^2): когда массив уже отсортирован или все элементы одинаковые
// На каждом уровне: n операций сравнения

func main() {
	fmt.Println("Простая сортировка")
	arr1 := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Исходный массив: %v\n", arr1)

	sorted1 := quickSort(arr1)
	fmt.Printf("Отсортированный: %v\n", sorted1)
}

// quickSort сортирует срез целых чисел используя алгоритм быстрой сортировки
// Опорный элемент - первый элемент среза
func quickSort(arr []int) []int {
	// Базовый случай: если массив содержит 0 или 1 элемент, он уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Создаем копию массива, чтобы не изменять исходный
	result := make([]int, len(arr))
	copy(result, arr)

	// Выбираем опорный элемент (первый элемент)
	pivot := result[0]

	// Создаем срезы для элементов меньше и больше опорного
	var less, greater []int

	// Разделяем элементы (начиная с индекса 1, так как pivot это result[0])
	for i := 1; i < len(result); i++ {
		if result[i] <= pivot {
			less = append(less, result[i]) //
		} else {
			greater = append(greater, result[i])
		}
	}

	// Рекурсивно сортируем части и объединяем результат
	sortedLess := quickSort(less)
	sortedGreater := quickSort(greater)

	// Объединяем результат: отсортированные меньшие + pivot + отсортированные большие
	var final []int
	final = append(final, sortedLess...)    // добавим отсортированные меньшие
	final = append(final, pivot)            // добавляем опорный эл
	final = append(final, sortedGreater...) // добавим отсортир большие

	return final // вернем  отсорт массив
}
