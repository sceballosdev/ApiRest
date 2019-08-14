package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sceballosdev/ApiRest/server/utils"
	"log"
)

func GetConnection() (*sql.DB, error) {
	// Open connection with cockroach db
	db, err := sql.Open(utils.DriverName, fmt.Sprintf(utils.UrlConnection, utils.UserDb, utils.DatabaseName));

	// Check for connection errors
	if err != nil {
		log.Fatal(utils.ErrorMessageConnection, err)
	}

	// Create the "items" table.
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS items (id INT PRIMARY KEY DEFAULT unique_rowid(), title STRING, post STRING)"); err != nil {
		log.Fatal(utils.ErrorTableCreation, err)
	}

	// Insert two rows into the "items" table.
	/*if _, err := db.Exec(
		"INSERT INTO items (title, post) VALUES ('Titulo 1', 'Post 1'), ('Titulo 2', 'Post 2')"); err != nil {
		log.Fatal(err)
	}*/

	return db, nil
}
