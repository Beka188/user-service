package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

var DB *sql.DB

func InitDb() error {
	var err error
	DB, err = sql.Open("sqlite3", "./foo.db?parseTime=true")
	if err != nil {
		return err
	}
	createTableSQL := `CREATE TABLE IF NOT EXISTS users
	(ID INTEGER PRIMARY KEY AUTOINCREMENT,
	 Name STRING, 
	 Email STRING UNIQUE NOT NUll,
	 Username STRING UNIQUE NOT NULL, 
	 Password String,
	 ProfilePicture STRING,
	 Bio STRING,
	 CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	 UpdatedAt DATETIME DEFAULT CURRENT_TIMESTAMP)`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}
