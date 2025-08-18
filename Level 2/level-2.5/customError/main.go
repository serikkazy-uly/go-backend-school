package main

// Раскрытый ответ в README
type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error

	err = test()
	if err != nil {
		println("error") // вывод
		return
	}
	println("ok")
}
