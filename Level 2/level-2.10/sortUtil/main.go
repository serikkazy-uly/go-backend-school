package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// SortOptions хранит все флаги сортировки
type SortOptions struct {
	column      int  // -k флаг для сортировки по столбцу
	numeric     bool // -n флаг для числовой сортировки
	reverse     bool // -r флаг для обратной сортировки
	unique      bool // -u флаг для уникальных строк
	month       bool // -M флаг для сортировки по месяцам
	ignoreBlank bool // -b флаг для игнорирования хвостовых пробелов
	check       bool // -c флаг для проверки отсортированности
	human       bool // -h флаг для человекочитаемых размеров
}

// Месяцы для сортировки по -M флагу
var months = map[string]int{
	"jan": 1, "feb": 2, "mar": 3, "apr": 4, "may": 5, "jun": 6,
	"jul": 7, "aug": 8, "sep": 9, "oct": 10, "nov": 11, "dec": 12,
}

func main() {
	var opts SortOptions

	// Парсинг флагов
	flag.IntVar(&opts.column, "k", 0, "Сортировка по столбцу N")
	flag.BoolVar(&opts.numeric, "n", false, "Числовая сортировка")
	flag.BoolVar(&opts.reverse, "r", false, "Обратная сортировка")
	flag.BoolVar(&opts.unique, "u", false, "Только уникальные строки")
	flag.BoolVar(&opts.month, "M", false, "Сортировка по названию месяца")
	flag.BoolVar(&opts.ignoreBlank, "b", false, "Игнорировать хвостовые пробелы")
	flag.BoolVar(&opts.check, "c", false, "Проверить отсортированность данных")
	flag.BoolVar(&opts.human, "h", false, "Сортировка с учётом суффиксов размеров")

	flag.Parse()

	var input io.Reader = os.Stdin

	// Если указан файл в аргументах
	if len(flag.Args()) > 0 {
		filename := flag.Args()[0]
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Ошибка открытия файла: %v", err)
		}
		defer file.Close()
		input = file
	}

	lines, err := readLines(input)
	if err != nil {
		log.Fatalf("Ошибка чтения данных: %v", err)
	}

	// Проверка отсортированности, если указан флаг -c
	if opts.check {
		if !isSorted(lines, opts) {
			fmt.Fprintln(os.Stderr, "Данные не отсортированы")
			os.Exit(1)
		}
		fmt.Println("Данные отсортированы")
		return
	}

	// Сортировка
	sortLines(lines, opts)

	// Удаление дубликатов, если указан флаг -u
	if opts.unique {
		lines = removeDuplicates(lines)
	}

	// Вывод результата
	for _, line := range lines {
		fmt.Println(line)
	}
}

// readLines читает все строки из input
func readLines(input io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// sortLines сортирует строки согласно указанным опциям
func sortLines(lines []string, opts SortOptions) {
	sort.Slice(lines, func(i, j int) bool {
		less := compareLess(lines[i], lines[j], opts)
		if opts.reverse {
			return !less
		}
		return less
	})
}

// compareLess сравнивает две строки согласно опциям сортировки
func compareLess(a, b string, opts SortOptions) bool {
	// Обработка игнорирования хвостовых пробелов
	if opts.ignoreBlank {
		a = strings.TrimRightFunc(a, unicode.IsSpace)
		b = strings.TrimRightFunc(b, unicode.IsSpace)
	}

	// Получение значений для сравнения
	valA := getCompareValue(a, opts)
	valB := getCompareValue(b, opts)

	// Числовое сравнение
	if opts.numeric {
		numA, errA := strconv.ParseFloat(valA, 64)
		numB, errB := strconv.ParseFloat(valB, 64)

		if errA == nil && errB == nil {
			return numA < numB
		}
		// Если не удалось распарсить как число, сравниваем как строки
	}

	// Сортировка по месяцам
	if opts.month {
		monthA := months[strings.ToLower(valA)]
		monthB := months[strings.ToLower(valB)]

		if monthA != 0 && monthB != 0 {
			return monthA < monthB
		}
		// Если не месяцы, сравниваем как строки
	}

	// Человеко-читаемые размеры
	if opts.human {
		sizeA := parseHumanSize(valA)
		sizeB := parseHumanSize(valB)

		if sizeA >= 0 && sizeB >= 0 {
			return sizeA < sizeB
		}
		// Если не размеры, сравниваем как строки
	}

	// Обычное строковое сравнение
	return valA < valB
}

// getCompareValue получает значение для сравнения в зависимости от опций
func getCompareValue(line string, opts SortOptions) string {
	if opts.column > 0 {
		// Разделение по табуляции (по умолчанию)
		fields := strings.Split(line, "\t")
		if opts.column <= len(fields) {
			return fields[opts.column-1]
		}
		// Если столбец не существует, возвращаем пустую строку
		return ""
	}
	return line
}

// parseHumanSize парсит человеко-читаемые размеры
func parseHumanSize(s string) float64 {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return -1
	}

	// Проверяем последний символ на суффикс
	lastChar := strings.ToLower(string(s[len(s)-1]))
	var multiplier float64 = 1
	var numStr string

	switch lastChar {
	case "k":
		multiplier = 1024
		numStr = s[:len(s)-1]
	case "m":
		multiplier = 1024 * 1024
		numStr = s[:len(s)-1]
	case "g":
		multiplier = 1024 * 1024 * 1024
		numStr = s[:len(s)-1]
	case "t":
		multiplier = 1024 * 1024 * 1024 * 1024
		numStr = s[:len(s)-1]
	default:
		numStr = s
	}

	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return -1
	}

	return num * multiplier
}

// removeDuplicates удаляет повторяющиеся строки
func removeDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}

	return result
}

// isSorted проверяет, отсортированы ли вобще строки
func isSorted(lines []string, opts SortOptions) bool {
	for i := 1; i < len(lines); i++ {
		less := compareLess(lines[i-1], lines[i], opts)
		if opts.reverse {
			less = !less
		}

		// Если предыдущий элемент больше текущего, тогда не отсортировано
		if !less && lines[i-1] != lines[i] {
			return false
		}
	}
	return true
}
