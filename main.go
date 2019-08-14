package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/sceballosdev/ApiRest/server/database"
	"github.com/sceballosdev/ApiRest/server/endpoints"
	"github.com/sceballosdev/ApiRest/server/utils"
	"net/http"
)

type Handler struct {
	chi.Router
	*sql.DB
	err error
}

func main() {
	test := Handler{}
	test.InitRouter()
}

func (handler *Handler) InitRouter() {
	handler.Router = chi.NewRouter()
	handler.DB, handler.err = database.GetConnection()

	// Get newsfeed
	handler.Router.Get(utils.EndpointNewsfeed, endpoints.NewsfeedGetAll(handler.DB))
	// Post newsfeed
	//handler.Router.Post(utils.EndpointNewsfeed, endpoints.NewsfeedInsert())

	http.ListenAndServe(utils.ApiPort, handler.Router)
}
