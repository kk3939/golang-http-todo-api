package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	Id      int `gorm:"primary_key"`
	Name    string
	Content string
}

type ToDos []Todo

var Db *gorm.DB

func connect(count int) {
	dsn := "test:password@tcp(golang-http-todo-api-db)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(count)
			return
		}
		panic(err)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	var todos ToDos
	result := Db.Find(&todos)
	if err := result.Error; err != nil {
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(&todos)
}

func main() {
	var count = 30
	connect(count)

	if t := Db.Migrator().HasTable(&Todo{}); t {
		fmt.Println("already Todo table exist. Seeds does not create table.")
	} else {
		if err := Db.Migrator().CreateTable(&Todo{}); err != nil {
			panic(err)
		}
		var todos ToDos
		for i := 1; i <= 10; i++ {
			todos = append(todos, Todo{
				Id:      i,
				Name:    fmt.Sprintf("name_%d", i),
				Content: fmt.Sprintf("content_%d", i),
			})
		}
		Db.Create(&todos)
		fmt.Println("create ten todo data!")
	}
	http.HandleFunc("/todos", getTodos)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
