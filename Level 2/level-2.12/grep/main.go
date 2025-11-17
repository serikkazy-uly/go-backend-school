package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Config содержит конфигурацию grep
type Config struct {
	After      int    // -A флаг: строки после совпадения
	Before     int    // -B флаг: строки до совпадения
	Context    int    // -C флаг: строки контекста (до и после)
	Count      bool   // -c флаг: только количество совпадений
	IgnoreCase bool   // -i флаг: игнорировать регистр
	Invert     bool   // -v флаг: инвертировать поиск
	Fixed      bool   // -F флаг: фиксированная строка (не regex)
	LineNum    bool   // -n флаг: выводить номера строк
	Pattern    string // Шаблон для поиска
	Files      []string
}

// SetContext устанавливает значения Before и After на основе Context
func (c *Config) SetContext() {
	if c.Context > 0 {
		c.After = c.Context
		c.Before = c.Context
	}
}

// GrepResult хранит результат grep
type GrepResult struct {
	Lines []LineResult
	Count int
}

// LineResult представляет одну строку результата
type LineResult struct {
	LineNum int
	Content string
	IsMatch bool // true если это совпадение, false если контекст
}

func main() {
	config := parseFlags()

	if config.Pattern == "" {
		fmt.Fprintln(os.Stderr, "Ошибка: не указан шаблон для поиска")
		os.Exit(1)
	}

	// Если файлы не указаны, читаем из STDIN
	if len(config.Files) == 0 {
		result, err := grep(os.Stdin, config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
			os.Exit(1)
		}
		printResult(result, config)
	} else {
		// Обработка файлов
		for _, filename := range config.Files {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка открытия файла %s: %v\n", filename, err)
				continue
			}

			result, err := grep(file, config)
			file.Close()

			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка обработки файла %s: %v\n", filename, err)
				continue
			}

			printResult(result, config)
		}
	}
}

func parseFlags() *Config {
	config := &Config{}

	flag.IntVar(&config.After, "A", 0, "вывести N строк после совпадения")
	flag.IntVar(&config.Before, "B", 0, "вывести N строк до совпадения")
	flag.IntVar(&config.Context, "C", 0, "вывести N строк контекста")
	flag.BoolVar(&config.Count, "c", false, "только количество совпадений")
	flag.BoolVar(&config.IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&config.Invert, "v", false, "инвертировать поиск")
	flag.BoolVar(&config.Fixed, "F", false, "фиксированная строка")
	flag.BoolVar(&config.LineNum, "n", false, "выводить номера строк")

	flag.Parse()

	// Если указан -C, применяем его к -A и -B
	if config.Context > 0 {
		config.After = config.Context
		config.Before = config.Context
	}

	// Оставшиеся аргументы: первый - шаблон, остальные - файлы
	args := flag.Args()
	if len(args) > 0 {
		config.Pattern = args[0]
		if len(args) > 1 {
			config.Files = args[1:]
		}
	}

	return config
}

func grep(reader io.Reader, config *Config) (*GrepResult, error) {
	scanner := bufio.NewScanner(reader)
	result := &GrepResult{
		Lines: make([]LineResult, 0),
		Count: 0,
	}

	var lines []string
	// lineNum := 0

	// Читаем все строки в память (для контекста)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Создаем matcher для поиска
	matcher, err := createMatcher(config.Pattern, config)
	if err != nil {
		return nil, err
	}

	// Отслеживаем какие строки уже добавлены (для избежания дубликатов)
	addedLines := make(map[int]bool)
	var matchedLines []int

	// Находим все совпадения
	for i, line := range lines {
		if matcher(line) {
			matchedLines = append(matchedLines, i)
			result.Count++
		}
	}

	// Если нужно только количество
	if config.Count {
		return result, nil
	}

	// Добавляем совпадения и контекст
	for _, matchIdx := range matchedLines {
		// Добавляем строки ДО совпадения
		beforeStart := matchIdx - config.Before
		if beforeStart < 0 {
			beforeStart = 0
		}
		for i := beforeStart; i < matchIdx; i++ {
			if !addedLines[i] {
				result.Lines = append(result.Lines, LineResult{
					LineNum: i + 1,
					Content: lines[i],
					IsMatch: false,
				})
				addedLines[i] = true
			}
		}

		// Добавляем само совпадение
		if !addedLines[matchIdx] {
			result.Lines = append(result.Lines, LineResult{
				LineNum: matchIdx + 1,
				Content: lines[matchIdx],
				IsMatch: true,
			})
			addedLines[matchIdx] = true
		}

		// Добавляем строки ПОСЛЕ совпадения
		afterEnd := matchIdx + config.After + 1
		if afterEnd > len(lines) {
			afterEnd = len(lines)
		}
		for i := matchIdx + 1; i < afterEnd; i++ {
			if !addedLines[i] {
				result.Lines = append(result.Lines, LineResult{
					LineNum: i + 1,
					Content: lines[i],
					IsMatch: false,
				})
				addedLines[i] = true
			}
		}
	}

	return result, nil
}

func createMatcher(pattern string, config *Config) (func(string) bool, error) {
	if config.Fixed {
		// Фиксированная строка
		searchPattern := pattern
		if config.IgnoreCase {
			searchPattern = strings.ToLower(pattern)
		}

		return func(line string) bool {
			searchLine := line
			if config.IgnoreCase {
				searchLine = strings.ToLower(line)
			}

			matched := strings.Contains(searchLine, searchPattern)
			if config.Invert {
				return !matched
			}
			return matched
		}, nil
	}

	// Регулярное выражение
	flags := ""
	if config.IgnoreCase {
		flags = "(?i)"
	}

	re, err := regexp.Compile(flags + pattern)
	if err != nil {
		return nil, fmt.Errorf("ошибка компиляции регулярного выражения: %w", err)
	}

	return func(line string) bool {
		matched := re.MatchString(line)
		if config.Invert {
			return !matched
		}
		return matched
	}, nil
}

func printResult(result *GrepResult, config *Config) {
	if config.Count {
		fmt.Println(result.Count)
		return
	}

	for _, line := range result.Lines {
		if config.LineNum {
			fmt.Printf("%d:%s\n", line.LineNum, line.Content)
		} else {
			fmt.Println(line.Content)
		}
	}
}
