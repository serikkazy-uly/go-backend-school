package main

import "fmt"

func UpdateBitPosition(value int64, index int, isSet bool) int64 {
	pattern := int64(1) << index

	if isSet {
		return value | pattern // Устанавливаем бит в 1,
	} else {
		return value &^ pattern // Устанавливаем бит в 0
	}
}

func main() {
	result := UpdateBitPosition(5, 1, false)
	fmt.Printf("Результат: %d\n", result) // 0101 -> 0101 (no change)

}
