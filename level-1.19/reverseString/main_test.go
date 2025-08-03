package main

import (
	"testing"
)

type testCase struct {
	name  string
	input string
	want  string
}

func TestReverseString(t *testing.T) {
	testCases := []testCase{
		// Базовые случаи
		{
			name:  "пустая строка",
			input: "",
			want:  "",
		},
		{
			name:  "простая ASCII строка",
			input: "hello",
			want:  "olleh",
		},
		{
			name:  "ASCII с числами",
			input: "abc123",
			want:  "321cba",
		},
		{
			name:  "ASCII палиндром",
			input: "racecar",
			want:  "racecar",
		},

		{
			name:  "Кириллица",
			input: "главрыба",
			want:  "абырвалг",
		},
		{
			name:  "английский и русский", // смешанная строка
			input: "Hello мир",
			want:  "рим olleH",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ReverseString(tc.input)
			if result != tc.want {
				t.Errorf("ReverseString(%q) = %q, хотели %q", tc.input, result, tc.want)
			}
		})
	}
}
