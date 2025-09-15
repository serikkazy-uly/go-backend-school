package main

import (
	"reflect"
	"testing"
)

func TestSortLines(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		opts     SortOptions
		expected []string
	}{
		{
			name:     "basic string sort",
			input:    []string{"c", "a", "b"},
			opts:     SortOptions{},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "reverse sort",
			input:    []string{"a", "b", "c"},
			opts:     SortOptions{reverse: true},
			expected: []string{"c", "b", "a"},
		},
		{
			name:     "numeric sort",
			input:    []string{"10", "2", "1"},
			opts:     SortOptions{numeric: true},
			expected: []string{"1", "2", "10"},
		},
		{
			name:     "numeric reverse sort",
			input:    []string{"1", "2", "10"},
			opts:     SortOptions{numeric: true, reverse: true},
			expected: []string{"10", "2", "1"},
		},
		{
			name:     "column sort",
			input:    []string{"a\t3", "b\t1", "c\t2"},
			opts:     SortOptions{column: 2, numeric: true},
			expected: []string{"b\t1", "c\t2", "a\t3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := make([]string, len(tt.input))
			copy(lines, tt.input)
			
			sortLines(lines, tt.opts)
			
			if !reflect.DeepEqual(lines, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, lines)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "no duplicates",
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "with duplicates",
			input:    []string{"a", "b", "a", "c", "b"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "all duplicates",
			input:    []string{"a", "a", "a"},
			expected: []string{"a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeDuplicates(tt.input)
			
			// Специальная обработка для пустых слайсов
			if len(tt.expected) == 0 && len(result) == 0 {
				return // оба пустые - успех
			}
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestParseHumanSize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{
			name:     "simple number",
			input:    "100",
			expected: 100,
		},
		{
			name:     "kilobytes",
			input:    "1k",
			expected: 1024,
		},
		{
			name:     "megabytes",
			input:    "1m",
			expected: 1024 * 1024,
		},
		{
			name:     "gigabytes",
			input:    "2g",
			expected: 2 * 1024 * 1024 * 1024,
		},
		{
			name:     "invalid input",
			input:    "abc",
			expected: -1,
		},
		{
			name:     "decimal with suffix",
			input:    "1.5k",
			expected: 1.5 * 1024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseHumanSize(tt.input)
			
			if result != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

func TestGetCompareValue(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		opts     SortOptions
		expected string
	}{
		{
			name:     "whole line",
			line:     "hello world",
			opts:     SortOptions{},
			expected: "hello world",
		},
		{
			name:     "first column",
			line:     "hello\tworld\tfoo",
			opts:     SortOptions{column: 1},
			expected: "hello",
		},
		{
			name:     "second column",
			line:     "hello\tworld\tfoo",
			opts:     SortOptions{column: 2},
			expected: "world",
		},
		{
			name:     "non-existent column",
			line:     "hello\tworld",
			opts:     SortOptions{column: 5},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getCompareValue(tt.line, tt.opts)
			
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestMonthSort(t *testing.T) {
	input := []string{"Dec", "Jan", "Feb", "Mar"}
	expected := []string{"Jan", "Feb", "Mar", "Dec"}
	
	opts := SortOptions{month: true}
	sortLines(input, opts)
	
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		opts     SortOptions
		expected bool
	}{
		{
			name:     "sorted ascending",
			lines:    []string{"a", "b", "c"},
			opts:     SortOptions{},
			expected: true,
		},
		{
			name:     "not sorted",
			lines:    []string{"c", "a", "b"},
			opts:     SortOptions{},
			expected: false,
		},
		{
			name:     "sorted descending with reverse flag",
			lines:    []string{"c", "b", "a"},
			opts:     SortOptions{reverse: true},
			expected: true,
		},
		{
			name:     "numeric sorted",
			lines:    []string{"1", "2", "10"},
			opts:     SortOptions{numeric: true},
			expected: true,
		},
		{
			name:     "numeric not sorted",
			lines:    []string{"10", "2", "1"},
			opts:     SortOptions{numeric: true},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSorted(tt.lines, tt.opts)
			
			if result != tt.expected {
				t.Errorf("Expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestIgnoreBlankSpaces(t *testing.T) {
	input := []string{"a  ", "b", "c   "}
	expected := []string{"a  ", "b", "c   "}
	
	opts := SortOptions{ignoreBlank: true}
	sortLines(input, opts)
	
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}