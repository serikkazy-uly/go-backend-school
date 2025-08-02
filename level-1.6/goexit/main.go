package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 1. Остановка с использованием контекста | самый применяемый способ
func method_Context() {
	fmt.Println("КОНТЕКСТ")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Запускаем 2 горутины
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Горутина %d остановлена\n", id)
					return
				default:
					fmt.Printf("Горутина %d работает\n", id)
					time.Sleep(500 * time.Millisecond)
				}
			}
		}(i)
	}

	// Останавливаем через 2 секунды
	time.Sleep(2 * time.Second)
	fmt.Println(">>> Отправляем cancel()")

	cancel()
	wg.Wait()
	fmt.Println("Завершено через контекст\n")
}

// 2: Использование канала-уведомления | классический способ
func method_Channel() {
	fmt.Println("КАНАЛ-УВЕДОМЛЕНИЯ")
	stopChan := make(chan struct{})
	var wg sync.WaitGroup

	// Запускаем горутины
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for {
				select {
				case <-stopChan:
					fmt.Printf("Горутина %d остановлена через канал\n", id)
					return
				default:
					fmt.Printf("Горутина %d работает\n", id)
					time.Sleep(400 * time.Millisecond)
				}
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(">>> Закрываем канал")

	close(stopChan)
	wg.Wait()

	fmt.Println("Завершено через канал-уведомления\n")
}

// 3: Использование sync/atomic флаг | производительность и безопасность
func method_Atomic() {
	fmt.Println("ATOMIC ФЛАГ")

	var shouldStop int32
	var wg sync.WaitGroup

	// Запускаем горутины
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for {
				if atomic.LoadInt32(&shouldStop) == 1 {
					fmt.Printf("Горутина %d остановлена через atomic флаг\n", id)
					return
				}

				fmt.Printf("Горутина %d работает\n", id)
				time.Sleep(300 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(">>> Устанавливаем atomic флаг")
	atomic.StoreInt32(&shouldStop, 1) // флаг остановки

	wg.Wait()
	fmt.Println("Завершено через atomic флаг\n")
}

// 4. Остановка через "mutex и фла"г или "Выход по условию" | базовая синхронизация
func method_Mutex() {
	fmt.Println("MUTEX + ФЛАГ")

	var stop bool // Флаг остановки
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Запускаем 2 горутины
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for {
				mu.Lock()
				shouldStop := stop
				mu.Unlock()

				if shouldStop {
					fmt.Printf("Горутина %d остановлена\n", id)
					return
				}

				fmt.Printf("Горутина %d работает\n", id)
				time.Sleep(500 * time.Millisecond)
			}
		}(i)
	}

	// Останавливаем через 2 секунды
	time.Sleep(2 * time.Second)
	fmt.Println(">>> Устанавливаем флаг через mutex")

	mu.Lock()
	stop = true
	mu.Unlock()

	wg.Wait()
	fmt.Println("Завершено через mutex\n")
}

// 5. Остановка через runtime.Goexit() | экстренные ситуации
func method_StopByGoexit() {
	fmt.Println("RUNTIME.GOEXIT()")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("Defer выполнен перед Goexit")

		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)

			if i == 2 {
				fmt.Println("КРИТИЧЕСКАЯ ОШИБКА - вызываем runtime.Goexit()") // Критическая ошибка - Экстренная остановка
				runtime.Goexit()                                              // Принудительно завершает горутину
			}
		}
	}()

	wg.Wait()
	fmt.Println("Завершено через runtime.Goexit()\n")
}

// 6. Остановка через таймер | для периодических задач (остановка по времени)
func method_Timer() {
	fmt.Println("ТАЙМЕР")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		timer := time.After(2 * time.Second) // Таймер на 2 секунды

		for {
			select {
			case <-timer:
				fmt.Println(">>>> Время истекло, горутина остановлена")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(400 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Завершено через таймер\n")
}

func main() {
	fmt.Println("Запуск всех способов остановки горутин")
	method_Context()
	method_Channel()
	method_Atomic()
	method_StopByGoexit()
	method_Timer()
	method_Mutex()
}
