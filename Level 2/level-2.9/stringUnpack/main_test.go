package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name:     "basic unpacking",
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
			hasError: false,
		},
		{
			name:     "no digits",
			input:    "abcd",
			expected: "abcd",
			hasError: false,
		},
		{
			name:     "only digits - error case",
			input:    "45",
			expected: "",
			hasError: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
			hasError: false,
		},
		{
			name:     "escape sequences",
			input:    "qwe\\4\\5",
			expected: "qwe45",
			hasError: false,
		},
		{
			name:     "partial escape",
			input:    "qwe\\45",
			expected: "qwe44444",
			hasError: false,
		},
		{
			name:     "single character with digit",
			input:    "a3",
			expected: "aaa",
			hasError: false,
		},
		{
			name:     "zero repetition",
			input:    "a0",
			expected: "",
			hasError: false,
		},
		{
			name:     "start with digit - error",
			input:    "4abc",
			expected: "",
			hasError: true,
		},
		{
			name:     "consecutive digits - error",
			input:    "a23",
			expected: "",
			hasError: true,
		},
		{
			name:     "escape backslash",
			input:    "a\\\\2",
			expected: "a\\\\",
			hasError: false,
		},
		{
			name:     "complex case with escapes",
			input:    "a2b\\34c5",
			expected: "aab3333ccccc",
			hasError: false,
		},
		{
			name:     "escape at end",
			input:    "a2\\3",
			expected: "aa3",
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := UnpackString(tt.input)
			
			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error for input %q, but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input %q: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("For input %q, expected %q, but got %q", tt.input, tt.expected, result)
				}
			}
		})
	}
}

func TestUnpackStringEdgeCases(t *testing.T) {
	// Дополнительные краевые случаи
	edgeCases := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name:     "single escape backslash",
			input:    "\\",
			expected: "\\",
			hasError: false, // обратный слеш в конце как символ
		},
		{
			name:     "multiple characters",
			input:    "ab2c3d4",
			expected: "abbcccdddd",
			hasError: false,
		},
		{
			name:     "unicode characters",
			input:    "ф2ы3",
			expected: "ффыыы",
			hasError: false,
		},
	}

	for _, tt := range edgeCases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := UnpackString(tt.input)
			
			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error for input %q, but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input %q: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("For input %q, expected %q, but got %q", tt.input, tt.expected, result)
				}
			}
		})
	}
}