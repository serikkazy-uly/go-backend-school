package main

import (
	"testing"
)

// false - сбросить бит, true - установить бит
func TestUpdateBitPosition(t *testing.T) {
	tests := []struct {
		value    int64
		index    int
		isSet    bool
		expected int64
	}{
		{5, 1, false, 5}, // 0101 -> 0101 (без изменений)
		{5, 1, true, 7},  // 0101 -> 0111 (установлен 1-й бит)
		{5, 2, false, 1}, // 0101 -> 0001 (сброшен 2-й бит)
		{5, 2, true, 5},  // 0101 -> 0101 (без изменений)
		{0, 0, true, 1},  // 0000 -> 0001 (установлен 0-й бит)
		{0, 0, false, 0}, // 0000 -> 0000 (без изменений)

	}

	for _, test := range tests {
		result := UpdateBitPosition(test.value, test.index, test.isSet)
		if result != test.expected {
			t.Errorf("UpdateBitPosition(%d, %d, %v) = %d; expected %d",
				test.value, test.index, test.isSet, result, test.expected)
		} else {
			t.Logf("UpdateBitPosition(%d, %d, %v) = %d; passed",
				test.value, test.index, test.isSet, result)
		}
	}
}
