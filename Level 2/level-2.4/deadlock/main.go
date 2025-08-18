package main

// close(ch) единственный способ сообщить range, что данных больше не будет
func main() {
	ch := make(chan int)

	go func() {
		// Deadlock - никто не пишет в канал, но range все еще ждет
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//Без close() range будет ждать бесконечно
		// close(ch) // добавить чтобы range корректно завершился
	}()
	for n := range ch {
		println(n)
	}
}

/*
Вывод:
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

Проблема:
Deadlock - никто больше не будет писать в канал, но range продолжает ждать
*/
