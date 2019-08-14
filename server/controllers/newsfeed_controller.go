package controllers

import (
	"database/sql"
	"fmt"
	"github.com/sceballosdev/ApiRest/server/models"
	"github.com/sceballosdev/ApiRest/server/utils"
	"log"
)

type Database struct {
	*sql.DB
}

func (db *Database) NewsfeedGetAll() ([]models.Item, error) {
	// Print out the balances.
	rows, err := db.Query(utils.NewsfeedGetAllQuery)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	items := []models.Item{}
	fmt.Println("Initial items:")

	for rows.Next() {
		item := models.Item{}
		if err := rows.Scan(&item.Id, &item.Title, &item.Post); err != nil {
			log.Fatal(err)
			return nil, err
		}

		items = append(items, item)
	}

	return items, err
}
