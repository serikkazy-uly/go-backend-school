# L2.3

## Что выведет программа? 
## Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
  "fmt"
  "os"
)
//происходит неявное преобразование типизированного nil в интерфейс error
func Foo() error {
  var err *os.PathError = nil // типизированный nil (тип *os.PathError)

  return err // возврат типизированный nil
}

func main() {
  err := Foo()
  fmt.Println(err)
  fmt.Println(err == nil)
}

```


Что получилось в наешм варианте:
```bash
      <nil>
      false
```


```go
// Внутри интерфейса error после возврата:
error {
    type:  *os.PathError  // тип сохранился
    value: nil            // значение nil
}
```
- В задании `type = *os.PathError`, поэтому `err != nil` даже если он пустой

- Интерфейс является nil в случае когда и тип(`type`) и тип значение (`value`) равны `nil`


- Интерфейсы это коллекции или набор методов (описывает поведение объекта и можно сказать глядя на методы)

- Внутреннее устройство интерфейсов:
```go

// обычный интерфейс
type iface struct {
    tab  *itab  // информация о типе | tab — это указатель на itable — структуру
    data unsafe.Pointer // указатель на данные
}

//Структура itab (Interface Table)
type itab struct {
    inter *interfacetype  // тип интерфейса (например err)
    _type *_type         // конкретный тип (наш вариант *os.PathError)
    hash  uint32         // копия _type.hash для быстрого сравнения
    _     [4]byte        // выравнивание
    
    // Таблица методов fun содержит указатели на конкретные реализации 
    // методов в порядке их объявления в интерфейсе
    fun   [1]uintptr     // таблица методов (variable sized)
}

// пустой тинтерфейс
type eface struct {
    _type *_type        //  только тип, без методов
    data  unsafe.Pointer // указатель на данные
}

```

- Итог нужно всегда возвращать явный nil для интерфейсов `return nil`
---


```go
// Интерфейсы и nil - дополнительные случаи

// Случай 1: nil интерфейс
var i interface{}  // i == nil (type: nil, data: nil)

// Случай 2: интерфейс с nil указателем - Наш случай из задания
var p *int = nil
var i interface{} = p  // i != nil (type: *int, data: nil)

// Случай 3: интерфейс с nil slice
var s []int = nil
var i interface{} = s  // i != nil (type: []int, data: nil)

```

- Интерфейс - это всегда структура из двух указателей:
```go
// Псевдокод внутреннего представления
type interface struct {
    type_or_itab *pointer  // 8 байт на 64-bit системе
    data         *pointer  // 8 байт на 64-bit системе
}    
// итого 16 байт
```