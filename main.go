package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Todos []Todo

func getTodo(w http.ResponseWriter, r *http.Request) {
	todos := Todos{}
	for i := 1; i <= 10; i++ {
		todos = append(todos, Todo{Id: i, Name: fmt.Sprintf("name_%d", i), Content: fmt.Sprintf("content_%d", i)})
	}
	json.NewEncoder(w).Encode(todos)
}

func main() {
	fmt.Println("Start....")
	http.HandleFunc("/todo", getTodo)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
