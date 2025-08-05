package main

import (
	"fmt"
	"time"
)

func customSleep(duration time.Duration) {
	// Создаем канал для получения сигнала о завершении ожидания
	done := make(chan bool)

	// Запускаем горутину, которая отправит сигнал через заданное время
	go func() {
		// time.After(duration) создает канал и внутренний таймер на 1 секунду
		<-time.After(duration) // блокируется на 1 секунду

		//  Через 1 сек таймер срабатывает и канал получает значение
		//  Горутина разблокируется и отправляет сигнал
		done <- true
	}()

	//Основная горутина ждет сигнала
	<-done // Блокируемся до получения true
}

func main() {
	fmt.Print("customSleep (1 сек)... ")
	start := time.Now()
	customSleep(1 * time.Second)

	fmt.Printf("Выполнено за: %v\n", time.Since(start))
}
