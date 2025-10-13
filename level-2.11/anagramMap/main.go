package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Тестовые данные из задания
	wordList := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	//поиск анаграмм
	anagrams := findAnagrams(wordList)
	fmt.Println(anagrams)
}

// Используется для определения, являются ли слова анаграммами
func sortRunes(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// findAnagrams находит все множества анаграмм в заданном словаре
// Возвращает map, где ключ - первое слово множества, значение - все анаграммы
func findAnagrams(words []string) map[string][]string {
	// Временный map для группировки слов по канонической форме
	anagramMap := make(map[string][]string)
	// Группируем слова по их канонической форме
	for _, word := range words {
		// Приводим к нижнему регистру
		normalizidWord := strings.ToLower(word)
		// Сортируем буквы в
		sorted := sortRunes(normalizidWord)
		anagramMap[sorted] = append(anagramMap[sorted], normalizidWord)
	}

	// Результирующий map
	result := make(map[string][]string)

	// Формируем результат выбирая только анаграммы
	for _, group := range anagramMap {
		// пропускаем группы из одного слова
		if len(group) < 2 {
			continue
		}
		// Сортируем по возрастанию
		sort.Strings(group)

		// Первое слово = ключ
		firstWord := group[0]
		result[firstWord] = group
	}
	return result
}
