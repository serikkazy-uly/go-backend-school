package main

import "fmt"

// Современный  логгер
// Клиент хочет простые методы LogInfo и LogError
type Logger interface {
	LogInfo(message string)  // логирование информационных сообщений
	LogError(message string) // логирование ошибок
}

// Существующая система логов - адаптируемый логер
type FileLogger struct {
	FileName string
}

// А старая система  с несовместимым интерфейсом (много ручной работы с логами)
func (f *FileLogger) WriteToFile(level, text string) {
	fmt.Printf("Запись в %s: [%s] %s\n", f.FileName, level, text)
}

// создаем Адаптер современный логгер
// FileLoggerAdapter - адаптер, который превращает старый FileLogger
// в совместимый с современным интерфейсом Logger
type FileLoggerAdapter struct {
	fileLogger *FileLogger
}

// метод адаптера
// клиент вызывает LogInfo(), а мы вызываем WriteToFile() с уровнем "INFO"
func (adapter *FileLoggerAdapter) LogInfo(message string) {
	adapter.fileLogger.WriteToFile("INFO", message)
}

// аналогично метод LogError()
// клиент вызывает LogError(), а мы вызываем WriteToFile() с уровнем "ERROR
func (adapter *FileLoggerAdapter) LogError(message string) {
	adapter.fileLogger.WriteToFile("ERROR", message)
}

func main() {
	fileLogger := &FileLogger{FileName: "app.log"}
	adapter := &FileLoggerAdapter{fileLogger: fileLogger}

	// Используем адаптер
	adapter.LogInfo("Информационное сообщение")
	adapter.LogError("Сообщение об ошибке")
}

/*
## Применимость паттерна

Из плюсов:
1) Принцип открытости/закрытости из SOLID:
Можем добавлять новые адаптеры без изменения клиентского кода:
var logger Logger = &FileLoggerAdapter{...}     // Файл
var logger Logger = &DatabaseLoggerAdapter{...} // БД

2)  Разделение ответственности:
- Бизнес-логика отделена от преобразования интерфейсов
- Адаптер занимается только адаптацией
- Клиент не знает о различиях в интерфейсах

Из минусов:

1) Усложнение кода
Дополнительный слой абстракции
Больше классов и интерфейсов
---
---

## Реальные примеры использования - Это то что везде есть скорее всего ну то что я точно встречал как миниммум

1) работа с БД

// Разные БД, единый интерфейс
type Database interface {
    Query(sql string) ([]Row, error)
}

type MySQLAdapter struct { client *mysql.Client }
type PostgreSQLAdapter struct { client *postgres.Client }
type MongoAdapter struct { client *mongo.Client } // NoSQL → SQL

// Приложение работает с любой БД одинаково
db := getDatabase() // MySQL, PostgreSQL или MongoDB
rows := db.Query("SELECT * FROM users")

2) Логирование - Мой пример

type Logger interface {
    LogInfo(message string)
    LogError(message string)
}

type LogrusAdapter struct { logrus *logrus.Logger }
type ZapAdapter struct { zap *zap.Logger }
type FileLoggerAdapter struct { file *FileLogger }

// Можно менять логгер без изменения кода
logger := getLogger()
logger.LogInfo("Application started")

3) Источники конфигурации
	...
*/
