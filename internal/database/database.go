package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Подключаем драйвер анонимно (через _)
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "./chat.db") // открываем базу, 1 аргумент - имя драйвера, второй - путь к файлу
	if err != nil {
		log.Fatal("ошибка подключения к базе данных:", err)
	}
	if err = db.Ping(); err != nil { // проверка что база работает посредством встроенного метода ping
		log.Fatal("База данных не отвечает")
	}
	query := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL
	);`
	msgQuery := `CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`
	// чуть выше, мы при помощи текстовой команды на языке SQL(язык общения с базами данных) - дали команды создать таблицы

	_, err = db.Exec(query) // собственно сам SQL запрос

	if err != nil {
		log.Fatal("Ошибка SQL запроса:", err)
	}
	_, err = db.Exec(msgQuery)

	if err != nil {
		log.Fatal("Ошибка SQL запроса:", err)
	}
	return db
}
