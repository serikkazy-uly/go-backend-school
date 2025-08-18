# L2.5

## Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
  msg string
}

func (e *customError) Error() string {
  return e.msg
}

func test() *customError {
  // ... do something
  return nil // нет ошибки
}

func main() {
  var err error
  err = test() // возвращает nil типа *customError
  if err != nil {
    println("error")
    return
  }
  println("ok")
}
```
---

Вывод: error

Что происходит:
- Функция `test()` возвращает `nil` - это `nil` типа `*customError`
- Присваивание `err = test()` - здесь nil указатель типа *customError 
присваивается переменной типа error (интерфейс)
---
После присваивания err содержит:
```text
Тип: *customError
Значение: nil
```

---
- Проверка `err != nil` - возвращает `true`
- Почему `err != nil` возвращает `true`?
- потому что в Go интерфейс состоит из 2-х частей:

```text
Тип (type)
Значение (value)
```
---
Интерфейс считается `nil` только тогда, когда и тип, и значение равны `nil`!!!

---
После выполнения err = test():
- err имеет тип: *customError
- err имеет значение: nil
---
Итог:
Поскольку тип не nil (это *customError), весь интерфейс не равен nil
