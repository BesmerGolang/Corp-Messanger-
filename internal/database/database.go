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
	// чуть выше, мы при помощи текстовой команды на языке SQL(язык общения с базами данных) - дали команду создать таблицу с полями id,username,password_hash
	_, err = db.Exec(query) // собственно сам SQL запрос
	if err != nil {
		log.Fatal("Ошибка SQL запроса:", err)
	}
	return db
}
