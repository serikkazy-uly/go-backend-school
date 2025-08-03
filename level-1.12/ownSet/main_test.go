package main

import (
	"reflect"
	"testing"
)

func TestOwnSet(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "Кейс из условия",
			input: []string{"cat", "dog", "cat", "tree"},
			want:  []string{"cat", "dog", "tree"},
		},
		{
			name:  "Без дубликатов",
			input: []string{"apple", "banana", "cherry"},
			want:  []string{"apple", "banana", "cherry"},
		},
		{
			name:  "Смешанный случай",
			input: []string{"apple", "banana", "apple", "cherry", "banana"},
			want:  []string{"apple", "banana", "cherry"},
		},
		{
			name:  "Все одинаковые",
			input: []string{"apple", "apple", "apple"},
			want:  []string{"apple"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OwnSet(tt.input)

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("OwnSet() = %v, want %v", result, tt.want)
			}
		})
	}
}
