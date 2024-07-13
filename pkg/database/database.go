package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
	"log"
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

	createTriggerSQL := `
	CREATE TRIGGER IF NOT EXISTS update_timestamp
	AFTER UPDATE ON users
	FOR EACH ROW
	BEGIN
		UPDATE users
		SET UpdatedAt = CURRENT_TIMESTAMP
		WHERE ID = OLD.ID;
	END;`

	_, err = DB.Exec(createTriggerSQL)
	if err != nil {
		log.Fatalf("Failed to create trigger: %v", err)
	}

	return nil
}
