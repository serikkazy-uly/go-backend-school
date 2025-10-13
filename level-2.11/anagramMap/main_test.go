package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "Пример из задания",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:     "Нет анаграмм",
			input:    []string{"дом", "кот", "лес"},
			expected: map[string][]string{},
		},
		{
			name:  "Все слова - анаграммы",
			input: []string{"кот", "ток", "кто"},
			expected: map[string][]string{
				"кот": {"кот", "кто", "ток"},
			},
		},
		{
			name:  "Разный регистр",
			input: []string{"Кот", "ТОК", "кто"},
			expected: map[string][]string{
				"кот": {"кот", "кто", "ток"},
			},
		},
		{
			name:     "Одно слово",
			input:    []string{"слово"},
			expected: map[string][]string{},
		},
		{
			name:     "Пустой список",
			input:    []string{},
			expected: map[string][]string{},
		},
		{
			name:  "Дубликаты слов",
			input: []string{"кот", "кот", "ток"},
			expected: map[string][]string{
				"кот": {"кот", "кот", "ток"},
			},
		},
		{
			name:  "Несколько групп анаграмм",
			input: []string{"сон", "нос", "кот", "ток", "дом"},
			expected: map[string][]string{
				"кот": {"кот", "ток"},
				"нос": {"нос", "сон"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagrams(tt.input)

			// Проверяем количество ключей
			if len(result) != len(tt.expected) {
				t.Errorf("Ожидалось %d групп анаграмм, получено %d", len(tt.expected), len(result))
			}

			// Проверяем каждую группу
			for key, expectedGroup := range tt.expected {
				resultGroup, exists := result[key]
				if !exists {
					t.Errorf("Ключ '%s' не найден в результате", key)
					continue
				}

				if !reflect.DeepEqual(resultGroup, expectedGroup) {
					t.Errorf("Для ключа '%s' ожидалось %v, получено %v", key, expectedGroup, resultGroup)
				}
			}
		})
	}
}

func TestSortRunes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Русские буквы",
			input:    "пятак",
			expected: "акптя",
		},
		{
			name:     "Уже отсортировано",
			input:    "абв",
			expected: "абв",
		},
		{
			name:     "Обратный порядок",
			input:    "яшу",
			expected: "ушя",
		},
		{
			name:     "Пустая строка",
			input:    "",
			expected: "",
		},
		{
			name:     "Одна буква",
			input:    "а",
			expected: "а",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortRunes(tt.input)
			if result != tt.expected {
				t.Errorf("sortRunes(%s) = %s, ожидалось %s", tt.input, result, tt.expected)
			}
		})
	}
}
