package main

import (
	"log"
	"net/http"

	"github.com/kk3939/golang-http-todo-api/database"
	"github.com/kk3939/golang-http-todo-api/server"
)

func main() {
	var rc = 30
	database.Connect(rc)
	database.Seeds()
	server.Router()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
