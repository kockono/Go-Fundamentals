package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Permite tener una variable global para la conexión a la base de datos
var DB *sql.DB

// DB is the database connection
func InitDb() { // *sql.DB
	_, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err)
	}

	// Maneja la cantidad de conexiones abiertas que se pueden tener
	DB.SetMaxOpenConns(10)
	// Maneja la cantidad de conexiones inactivas que se pueden tener
	DB.SetMaxIdleConns(5)
	createTables()
	// return db
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	// Ejecuta la query
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}

}
