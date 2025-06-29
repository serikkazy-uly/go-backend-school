package main

import "fmt"

type Human struct {
	Name       string
	Age        int
	Profession string
}

// Теперь Action автоматически "наследует" поля и методы Human
type Action struct {
	Work  string
	Relax string
	Human
}

func (h Human) Walk() string {
	return fmt.Sprintf("%s часто гуляет пешком, потому что прогулка приносит пользу для профессии %s", h.Name, h.Profession)

}

func (h Human) Run() string {
	return fmt.Sprintf("%s бегает каждый день,  потому что любит оставаться в бодрости в свои %d лет.", h.Name, h.Age)
}

func (a Action) Do() string {
	// Здесь мы можем использовать поля и методы встраиваемой структуры Human
	return fmt.Sprintf("%s кроме того, по профессии %s но по вечерам  %s.", a.Name, a.Profession, a.Relax)
}

func main() {
	h := Human{
		Name:       "Дамир",
		Age:        36,
		Profession: "Инженер-программист",
	}

	fmt.Println("--- Пример использования структуры Human ---")
	// Создаем экземпляр структуры Action.
	// При инициализации Action мы инициализируем и встроенную структуру Human.
	a := Action{
		Human: Human{"Алиса", 28, "бухгалтер"}, // Встраиваем Human
		Work:  "программирование",
		Relax: "ходит в зал",
	}

	fmt.Println(h.Walk()) // Вызов метода Human
	fmt.Println(h.Run())  // Вызов метода Human

	fmt.Println("\n--- Пример встраивания Human в Action ---")
	fmt.Println(a.Walk()) // Вызов метода напрямую поля из Human через переменную 'a'
	fmt.Println(a.Run())  // Вызов метода напрямую поля из Human через переменную 'a'

	fmt.Println(a.Do()) // Вызываем собственный метод Action
	// Доступ к полю Name из встраиваемого Human напрямую через Action
	fmt.Printf("Имя (из Action): %s\n", a.Name)
}
