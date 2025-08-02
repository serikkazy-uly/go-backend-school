package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	safeMutex()
}

// использовал sync.Mutex на блокировках
func safeMutex() {
	var (
		safeMap = make(map[string]int) // Инициализируем безопасную map
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	// 100 горутин безопасно пишут в map
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			key := fmt.Sprintf("worker_%d", id%10) // только 10 ключей для 100 горутин

			mu.Lock()              // Блокируем доступ - все горутины будут ждать своей очереди
			safeMap[key] = id * 10 // Безопасная запись
			mu.Unlock()            // Разблокируем

			time.Sleep(1 * time.Millisecond)
		}(i)
	}

	wg.Wait()

	// Безопасное чтение
	mu.Lock()
	fmt.Printf("Результат: %v\n", safeMap)
	mu.Unlock()

	fmt.Println("Все данные записаны безопасно")
}

// Терминал:
// go run -race main.go

// Вывод программы:
// Результат: map[worker_0:300 worker_1:710 worker_2:720 worker_3:930 worker_4:740 worker_5:850 worker_6:960 worker_7:970 worker_8:280 worker_9:990]
// Все данные записаны безопасно
