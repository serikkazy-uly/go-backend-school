package main

import (
	"fmt"
	"sort"
)

func GroupTemperaturesByStep(temperatures []float64, step int) map[int][]float64 {
	groups := make(map[int][]float64) // Создаем карту для групп

	for _, temp := range temperatures {

		group := int(temp) / step // Находим номер группы
		key := group * step       // Вычисляем ключ группы

		// Добавляем температуру в соответствующую группу
		groups[key] = append(groups[key], temp)
	}

	return groups
}

func main() {
	// Исходная последовательность температур
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println("Исходные температуры:", temperatures)
	fmt.Println()

	step := 10

	groups := GroupTemperaturesByStep(temperatures, step)
	fmt.Printf("Группы температур с шагом %d:\n", step)

	// Сортируем ключи групп для упорядоченного вывода
	keys := make([]int, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Выводим отсортированный результат
	for _, key := range keys {
		fmt.Printf("Группа %d: %v\n", key, groups[key])
	}
	// // Вывод результатов
	// for key, temps := range groups {
	// 	fmt.Printf("Группа %d: %v\n", key, temps)
	// }

}
