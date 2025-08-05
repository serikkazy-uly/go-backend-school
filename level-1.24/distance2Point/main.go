package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64 // приватные поля (начинаются с маленькой буквы)
}

// NewPoint - конструктор для создания новой точки
// Указатели эффективнее для больших структур, но для моего решения не критично
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Distance вычисляет расстояние от текущей точки до другой точки
// Передача по указателю избегает копирования
func (p *Point) Distance(other *Point) float64 {
	// По заданной Формуле расстояния: sqrt((x2-x1)^2 + (y2-y1)^2)
	dx := other.x - p.x
	dy := other.y - p.y

	return math.Sqrt(dx*dx + dy*dy)
}

// GetX возвращает координату x (геттер для приватного поля)
func (p *Point) GetX() float64 {
	return p.x
}

// GetY возвращает координату y (геттер для приватного поля)
func (p *Point) GetY() float64 {
	return p.y
}

// String возвращает строковое представление точки
func (p *Point) String() string {
	return fmt.Sprintf("Point(%.2f, %.2f)", p.x, p.y)
}

func main() {
	// Создаем две точки с помощью конструктора
	point1 := NewPoint(0.0, 0.0)   // Начало координат
	point2 := NewPoint(-1.0, -1.0) // Отрицательные координаты
	point3 := NewPoint(5.5, 2.3)   // Дробные координаты

	// Выводим информацию о точках
	fmt.Println("Точка 1:", point1)
	fmt.Println("Точка 2:", point2)
	fmt.Println("Точка 3:", point3)
	fmt.Println()

	// Вычисляем расстояния между разными точками
	distance1_2 := point1.Distance(point2) // от point1 до point2
	distance1_3 := point1.Distance(point3) // от point1 до point3
	distance2_3 := point2.Distance(point3) // от point2 до point3

	// Выводим результаты
	fmt.Printf("Расстояние между %s и %s: %.2f\n", point1, point2, distance1_2)
	fmt.Printf("Расстояние между %s и %s: %.2f\n", point1, point3, distance1_3)
	fmt.Printf("Расстояние между %s и %s: %.2f\n", point2, point3, distance2_3)
}
