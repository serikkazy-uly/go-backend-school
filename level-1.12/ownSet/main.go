package main

import (
	"fmt"
	"sort"
)

func OwnSet(mySet []string) []string {
	// множество для хранения уникальных элементов из A
	uniqSet := make(map[string]struct{})

	// заполняем множество эл-ми из mySet
	for _, item := range mySet {
		uniqSet[item] = struct{}{} //  struct весит - 0 кБ
	}

	// преобразем множество в срез
	// используем make с начальной емкостью равной кол-ву уник эл-ов
	resultOutput := make([]string, 0, len(uniqSet))
	// добавляем эл-ты из множества в срез
	// перебираем "ключи" множества
	for item := range uniqSet {
		resultOutput = append(resultOutput, item)
	}

	sort.Strings(resultOutput)
	return resultOutput
}

func main() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("Входные данные: %v\n", set)

	uniqueWords := OwnSet(set)
	fmt.Println("\nВыходные данные: ", uniqueWords)
}
