package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем время с NTP сервера
	time, err := ntp.Time("time.google.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения времени через NTP: %v\n", err)
		os.Exit(1)
	}

	// Выводим точное время
	fmt.Printf("Точное время (NTP): %s\n", time.Format("2006-01-02 15:04:05 MST"))
}