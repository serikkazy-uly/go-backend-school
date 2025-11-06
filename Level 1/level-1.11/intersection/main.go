package main

import (
	"fmt"
	"sort"
)

func Intersection(A, B []int) []int {
	// множество для хранения уникальных элементов из A
	setA := make(map[int]struct{})
	// заполняем множество элементами из A
	// используем пустую структуру для экономии памяти
	// ключи множества будут уникальными элементами из A
	// это позволяет быстро проверять наличие элемента в A
	for _, a := range A {
		setA[a] = struct{}{}
	}

	// множество для хранения общих элементов А и В
	commonElements := make(map[int]struct{})
	// перебираем элементы из B
	for _, b := range B {
		// если элемент из B есть в A, добавляем его в множество общих элементов
		// используем пустую структуру для экономии памяти
		if _, exists := setA[b]; exists {
			commonElements[b] = struct{}{}
		}
	}
	// преобразуем множество общих элементов в срез
	// используем make с начальной емкостью, равной количеству общих элементов
	result := make([]int, 0, len(commonElements))
	// добавляем элементы из множества в срез
	for elem := range commonElements {
		result = append(result, elem)
	}
	// сортируем результат для удобства сравнения
	sort.Ints(result)
	// возвращаем отсортированный срез общих элементов
	return result
}

func main() {
	// имеются повторяющиеся и неосортерованные элементы в A и B
	A := []int{4, 9, 2, 3, 3, 4, 4, 5}
	B := []int{9, 2, 3, 4, 5, 6, 6, 6}

	result := Intersection(A, B)

	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}
