package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// Permite tener una variable global para la conexión a la base de datos
var DB *sql.DB

// DB is the database connection
func InitDb() { // *sql.DB
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	// Maneja la cantidad de conexiones abiertas que se pueden tener
	DB.SetMaxOpenConns(10)
	// Maneja la cantidad de conexiones inactivas que se pueden tener
	DB.SetMaxIdleConns(10)
	createTables()
	// return db
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		password TEXT NOT NULL
		date_time DATETIME NOT NULL
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		FORIEGN KEY (user_id) REFERENCES users(id)
	)
	`
	// Ejecuta la query
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic(err)
	}

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
