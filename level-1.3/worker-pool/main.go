package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

/*
flow:
- работа по таймеру,
- Нет внешних сигналов (Ctrl+C)
- альтернатива Graceful shutdown через последовательное завершение
- Генератор пишет данные в канал jobs
- N воркеров конкурентно читают из канала jobs
- Каждый воркер выводит полученные данные в stdout через fmt.Printf
- Go runtime автоматически распределяет сообщения между доступными воркерами
*/

// читает из канала jobs и выводит их в stdout
func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Цикл завершится автоматически, когда канал-jobs будет закрыт и пуст
	for job := range jobs { // Чтение данных из канала - jobs
		fmt.Printf("Воркер %d начал задание: %s\n", id, job) // вывод в stdout
		time.Sleep(time.Second)                              // Имитация полезной работы
	}
	fmt.Printf("Воркер %d завершил работу\n", id)

}

// постоянно генерирует и записывает данные в канал
func dataProducer(jobs chan<- string, stop <-chan struct{}) {
	defer fmt.Println("Генератор данных завершен")

	counter := 1
	ticker := time.NewTicker(300 * time.Millisecond) // интервал для генератора
	defer ticker.Stop()
	fmt.Println("Генератор данных запущен")

	for {
		select {
		// 1-ый сценарий - остановка
		case <-stop:
			fmt.Println("Генератор получил сигнал остановки")
			return
		// 2-ой сценарий - В канал отправляется time.Time - момент когда должен был сработать тик
		case <-ticker.C: // Ждем сигнала от ticker'а каждые 300 ms
			data := fmt.Sprintf("Сообщение #%d", counter)
			// 3-ий сценарий защищает от блокировки при отправки
			select {
			// 1-ый сценарий
			case jobs <- data:
				fmt.Printf("Отправлено успешно: %s\n", data)
				counter++
			// 2-ой сценарий
			case <-stop:
				fmt.Println("Генератор остановлен во время отправки")
				return
			}
		}
	}
}

func cleanup(stop chan struct{}, jobs chan string) {
	close(stop) //Сначала останавливаем генератор
	log.Println("Generator остановлен")

	time.Sleep(100 * time.Millisecond)

	close(jobs) // безопасно закрываем канал jobs
	log.Println("Jobs канал закрыт")
}

// Создание N воркеров
func main() {
	// --- Получаем N воркеров из аргументов командной строки ---
	nWorkers := 3 // Значение по умолчанию

	//--- Валидация ---
	if len(os.Args) > 1 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil && n > 0 {
			nWorkers = n // пользователь задает N
		} else {
			fmt.Println("значение по умолчанию: 3")
		}
	}
	fmt.Printf("Запуск с %d воркерами...\n", nWorkers)

	// --- Инициализация ---
	// Создаем канал для передачи данных
	// Буферизованный канал для избежания блокировок
	jobs := make(chan string, 10) // Канал для данных
	stop := make(chan struct{})   // Канал для остановки генератора
	var wg sync.WaitGroup

	// --- Создаем именно N воркеров
	// Запускаем указанное количество - N воркеров
	for ws := 1; ws <= nWorkers; ws++ { // от 1 до N
		wg.Add(1)
		// Запускаем N горутин
		go worker(ws, jobs, &wg)
	}

	// Запускаем генератор данных (постоянно пишет в канал) в главной горутине-main
	go dataProducer(jobs, stop)

	// Работаем определенное время (для демонстрации)
	time.Sleep(10 * time.Second)

	cleanup(stop, jobs) //Блокируе main до завершения

	wg.Wait()
	fmt.Println("Все воркеры завершили свою работу. Программа завершена.")
}
