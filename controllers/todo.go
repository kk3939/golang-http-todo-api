package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kk3939/golang-http-todo-api/database"
	"github.com/kk3939/golang-http-todo-api/models"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTodos(w, r)
	case "POST":
		postTodo(w, r)
	case "PUT":
		updateTodo(w, r)
	case "DELETE":
		deleteTodo(w, r)
	default:
		w.WriteHeader(405)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	var todos models.ToDos
	result := database.Db.Find(&todos)
	if err := result.Error; err != nil {
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(&todos)
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		t := &models.Todo{
			Name:    todo.Name,
			Content: todo.Content,
		}
		database.Db.Create(t)
		json.NewEncoder(w).Encode(t)
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		database.Db.Model(&todo).Updates(models.Todo{Name: todo.Name, Content: todo.Content})
		json.NewEncoder(w).Encode(todo)
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		database.Db.Delete(&todo)
		json.NewEncoder(w).Encode(todo)
	}
}
