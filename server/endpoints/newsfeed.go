package endpoints

import (
	"database/sql"
	"encoding/json"
	"github.com/sceballosdev/ApiRest/server/controllers"
	"github.com/sceballosdev/ApiRest/server/models"
	"net/http"
)

func NewsfeedGetAll(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		controller := controllers.Database{DB: db}
		items, err := controller.NewsfeedGetAll()

		if err != nil {
			json.NewEncoder(writer).Encode("Error obteniendo datos " + err.Error())
			return
		}

		response := map[string][]models.Item{}
		response["items"] = items

		json.NewEncoder(writer).Encode(response)
	}
}

func NewsfeedInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		requestMap := map[string]string{}
		json.NewDecoder(request.Body).Decode(&requestMap)

		item := models.Item{
			Title: requestMap["title"],
			Post:  requestMap["post"],
		}

		writer.Write([]byte("Item was added " + item.Title))
	}
}
