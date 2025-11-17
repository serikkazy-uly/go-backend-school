package main

import (
	"strings"
	"testing"
)

func TestGrepBasic(t *testing.T) {
	input := `line one
line two
line three
line four`

	config := &Config{
		Pattern: "two",
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if result.Count != 1 {
		t.Errorf("Ожидалось 1 совпадение, получено %d", result.Count)
	}

	if len(result.Lines) != 1 {
		t.Errorf("Ожидалась 1 строка, получено %d", len(result.Lines))
	}

	if result.Lines[0].Content != "line two" {
		t.Errorf("Ожидалось 'line two', получено '%s'", result.Lines[0].Content)
	}
}

func TestGrepIgnoreCase(t *testing.T) {
	input := `Line One
			  LINE TWO
			  line three`

	config := &Config{
		Pattern:    "line",
		IgnoreCase: true,
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if result.Count != 3 {
		t.Errorf("Ожидалось 3 совпадения, получено %d", result.Count)
	}
}

func TestGrepInvert(t *testing.T) {
	input := "line one\nline two\ntext three\nline four"

	config := &Config{
		Pattern: "line",
		Invert:  true,
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if result.Count != 1 {
		t.Errorf("Ожидалось 1 совпадение, получено %d", result.Count)
	}

	if result.Lines[0].Content != "text three" {
		t.Errorf("Ожидалось 'text three', получено '%s'", result.Lines[0].Content)
	}
}

func TestGrepCount(t *testing.T) {
	input := `line one
			  line two
			  text three
			  line four`

	config := &Config{
		Pattern: "line",
		Count:   true,
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if result.Count != 3 {
		t.Errorf("Ожидалось 3 совпадения, получено %d", result.Count)
	}

	if len(result.Lines) != 0 {
		t.Errorf("При флаге -c не должно быть строк в результате, получено %d", len(result.Lines))
	}
}

func TestGrepAfter(t *testing.T) {
	input := `line 1
			  line 2
			  match
			  line 4
			  line 5`

	config := &Config{
		Pattern: "match",
		After:   2,
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// Ожидаем: match, line 4, line 5
	if len(result.Lines) != 3 {
		t.Errorf("Ожидалось 3 строки (совпадение + 2 после), получено %d", len(result.Lines))
	}
}

func TestGrepBefore(t *testing.T) {
	input := `line 1
			  line 2
			  match
			  line 4
			  line 5`

	config := &Config{
		Pattern: "match",
		Before:  2,
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// Ожидаем: line 1, line 2, match
	if len(result.Lines) != 3 {
		t.Errorf("Ожидалось 3 строки (2 до + совпадение), получено %d", len(result.Lines))
	}
}

func TestGrepContextDebug(t *testing.T) {
	input := `line 1
			  line 2
			  match
			  line 4
			  line 5`

	config := &Config{
		Pattern: "match",
		Context: 1,
	}

	config.SetContext()

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	t.Log("=== ДЕТАЛЬНАЯ ОТЛАДОЧНАЯ ИНФОРМАЦИЯ ===")
	t.Logf("Config: Pattern='%s', Context=%d", config.Pattern, config.Context)
	t.Logf("Total matches: %d", result.Count)
	t.Logf("Total lines in result: %d", len(result.Lines))

	for i, line := range result.Lines {
		marker := " "
		if line.IsMatch {
			marker = "★"
		}
		t.Logf("%s %2d: [%d] %s", marker, i, line.LineNum, line.Content)
	}

	// Проверяем конкретные ожидания
	if len(result.Lines) < 3 {
		t.Errorf("СЛИШКОМ МАЛО СТРОК: ожидалось 3, получено %d", len(result.Lines))
	} else {
		t.Log("✓ Количество строк соответствует ожиданиям")
	}
}

func TestGrepLineNum(t *testing.T) {
	input := `line 1
			  line 2
			  match
			  line 4`

	config := &Config{
		Pattern: "match",
		LineNum: true,
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if len(result.Lines) != 1 {
		t.Errorf("Ожидалась 1 строка, получено %d", len(result.Lines))
	}

	if result.Lines[0].LineNum != 3 {
		t.Errorf("Ожидался номер строки 3, получено %d", result.Lines[0].LineNum)
	}
}

func TestGrepFixed(t *testing.T) {
	input := `test.string
			  test*string
			  test[string]`

	// С флагом -F точка и звездочка должны восприниматься как литералы
	config := &Config{
		Pattern: "test.",
		Fixed:   true,
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if result.Count != 1 {
		t.Errorf("Ожидалось 1 совпадение (только с точкой), получено %d", result.Count)
	}
}

func TestGrepRegex(t *testing.T) {
	input := `test123
			  test456
			  testabc`

	config := &Config{
		Pattern: `test\d+`,
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if result.Count != 2 {
		t.Errorf("Ожидалось 2 совпадения, получено %d", result.Count)
	}
}

func TestGrepContextBoundary(t *testing.T) {
	// Тест граничного случая: совпадение в начале файла
	input := `match
			  line 2
			  line 3`

	config := &Config{
		Pattern: "match",
		Before:  5, // Больше чем доступно
		After:   1,
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// Ожидаем: match, line 2
	if len(result.Lines) != 2 {
		t.Errorf("Ожидалось 2 строки, получено %d", len(result.Lines))
	}
}

func TestGrepContextBoundaryEnd(t *testing.T) {
	// Тест граничного случая: совпадение в конце файла
	input := `line 1 
	          line 2 
			  match`

	config := &Config{
		Pattern: "match",
		Before:  1,
		After:   5, // Больше чем доступно
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// Ожидаем: line 2, match
	if len(result.Lines) != 2 {
		t.Errorf("Ожидалось 2 строки, получено %d", len(result.Lines))
	}
}

func TestGrepMultipleMatches(t *testing.T) {
	input := `match 1
				line 2
				match 3
				line 4
				match 5`

	config := &Config{
		Pattern: "match",
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if result.Count != 3 {
		t.Errorf("Ожидалось 3 совпадения, получено %d", result.Count)
	}

	if len(result.Lines) != 3 {
		t.Errorf("Ожидалось 3 строки, получено %d", len(result.Lines))
	}
}

func TestGrepOverlappingContext(t *testing.T) {
	// Тест случая с перекрывающимся контекстом
	input := `line 1
				match 2
				match 3
				line 4`

	config := &Config{
		Pattern: "match",
		Context: 1,
	}

	config.SetContext()

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// line 1, match 2, match 3, line 4 - без дубликатов
	if len(result.Lines) != 4 {
		t.Errorf("Ожидалось 4 уникальных строки, получено %d", len(result.Lines))
	}
}

func TestGrepCombinedFlags(t *testing.T) {
	input := `Line 1
				MATCH 2
				Line 3
				Line 4`

	config := &Config{
		Pattern:    "match",
		IgnoreCase: true,
		LineNum:    true,
		Context:    1,
	}

	config.SetContext()

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// Ожидаем Line 1, MATCH 2, Line 3
	if len(result.Lines) != 3 {
		t.Errorf("Ожидалось 3 строки, получено %d", len(result.Lines))
	}

	// Проверяем номера строк
	if result.Lines[0].LineNum != 1 || result.Lines[1].LineNum != 2 || result.Lines[2].LineNum != 3 {
		t.Errorf("Неправильные номера строк")
	}
}

func TestGrepEmptyInput(t *testing.T) {
	input := ``

	config := &Config{
		Pattern: "test",
	}

	reader := strings.NewReader(input)
	result, err := grep(reader, config)

	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if result.Count != 0 {
		t.Errorf("Ожидалось 0 совпадений, получено %d", result.Count)
	}
}
