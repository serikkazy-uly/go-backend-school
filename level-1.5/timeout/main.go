package main

import (
	"fmt"
	"time"
)

func main() {
	bufer := 2
	ch := make(chan int, bufer)
	timeout := time.After(5 * time.Second)

	// Producer: отправляет последовательные значения в канал
	go func() {
		defer close(ch)
		i := 1
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-timeout:
				fmt.Println("Producer: таймоут достигнут")
				return
			case <-ticker.C:
				select {
				case ch <- i:
					fmt.Printf("Sent: %d\n", i)
					i++
				case <-timeout:
					fmt.Println("Producer: таймоут на время отправки")
					return
				}
			default:
				ch <- i
				i++
				time.Sleep(500 * time.Millisecond) // немного задерживаем для имитации
			}
		}
	}()

	// Consumer: читает значения из канала
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт - выход")
				return
			}
			fmt.Println("Получен:", v)
		case <-timeout:
			fmt.Println("Timeout достигнут -выход")
			return
		}
	}
}
