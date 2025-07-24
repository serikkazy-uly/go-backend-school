package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Коротко: сигнал получается сразу после закрытия канала при cancel()
==================================================================
Выбрал  signal.NotifyContext по рекомендации старших инженеров коллег
- имеет автоматическую отмену по SIGINT во все горутины
- и эта функция из под коробки go
- go самостоятельно свяжет сигнал с контекстом
- ctx автоматически отменится при Ctrl+C
Как-буто исключает человеческий фактор (все под капотом реализовано) - забыл остановить, закрыть, недописал .. и тп
(согласен что неявно, но всегда можно провалитьтся в функцию и посмотреть что там)
*/
func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		//Все воркеры получили сигнал через <-ctx.Done() одновременно
		case <-ctx.Done(): // Реагируем на отмену контекста cancel()
			fmt.Printf("Worker %d: shutting down...\n", id)
			return
		default:
			fmt.Printf("Worker %d: работает...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// создаю ctx который отменяется при получении SIGINT(Ctrl+c)
	// Автоматическая связь сигнала с контекстом
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup
	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	fmt.Println("Нажмите Ctrl+C для остановки...")

	// Ожидание сигнала и graceful shutdown
	<-ctx.Done()
	fmt.Println("Main: получил SIGINT, shutting down...")

	wg.Wait()

}
