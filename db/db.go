package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTable()
}

func CreateTable() {

	createUserTable :=
		`CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create users DB")
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(createEventTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create event DB")
	}

}
