package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name string
		A, B []int
		want []int
	}{
		{
			name: "Test Case 1",
			A:    []int{1, 2, 3, 4, 5},
			B:    []int{2, 3, 4, 5, 6},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "Test Case 2", // нет пересечений
			A:    []int{7, 8, 9},
			B:    []int{10, 11, 12},
			want: []int{},
		},
		{
			name: "Test Case 3",
			A:    []int{1, 1, 2, 2, 3, 3},
			B:    []int{2, 2, 3, 3, 4, 4},
			want: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersection(tt.A, tt.B)

			sort.Ints(result)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Intersection() = %v, want %v", result, tt.want)
			}
		})
	}
}

// func TestIntersection(t *testing.T) {
// 	A := []int{111, 22, 13, 4, 5}
// 	B := []int{3, 4, 111, 22, 5, 13, 6, 7, 5}

// 	expected := []int{4, 5, 111, 22, 13}
// 	result := Intersection(A, B)

// 	sort.Ints(result)
// 	sort.Ints(expected)

// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("Expected %v, got %v", expected, result)
// 	}
// }
