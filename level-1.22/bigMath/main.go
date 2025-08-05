package main

import (
	"fmt"
	"math/big"
)

// calculate выполняет арифметические операции для больших чисел
func calculate(x, y *big.Float) {
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)

	// Массив для хранения результатов
	operations := []*big.Float{
		new(big.Float).Add(x, y), // сумма
		new(big.Float).Sub(x, y), // разность
		new(big.Float).Mul(x, y), // произведение
		new(big.Float).Quo(x, y), // частное
	}

	symbols := []string{" + ", " - ", " × ", " ÷ "}

	for i, result := range operations {
		fmt.Printf("x%sy = %v\n", symbols[i], result)
	}
	fmt.Println()
}

// createBigNumber создает большое число из строки
func createBigNumber(value string) *big.Float {
	num := new(big.Float)
	num.SetString(value)

	return num
}

func main() {
	fmt.Println("Арифметика больших чисел (значения > 2^20)")

	// Вариант 1: обычные большие числа
	a1 := createBigNumber("8888888")
	b1 := createBigNumber("7777777")
	calculate(a1, b1)

	// Вариант 2: числа с дробной частью
	a2 := createBigNumber("12345678901234567890.123")
	b2 := createBigNumber("98765432109876543210.456")
	calculate(a2, b2)

	// Вариант 3: экстремально большие числа
	a3 := createBigNumber("999999999999999999999999999999999999")
	b3 := createBigNumber("123456789012345678901234567890123456")
	calculate(a3, b3)
}
