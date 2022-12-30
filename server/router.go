package server

import (
	"net/http"

	"github.com/kk3939/golang-http-todo-api/controllers"
)

func Router() {
	http.HandleFunc("/todos", controllers.TodoHandler)
}
