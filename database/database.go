package database

import (
	"fmt"
	"time"

	"github.com/kk3939/golang-http-todo-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(count int) {
	dsn := "test:password@tcp(golang-http-todo-api-db)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			Connect(count)
			return
		}
		panic(err)
	}
}

func Seeds() {
	if t := Db.Migrator().HasTable(&models.Todo{}); t {
		fmt.Println("already Todo table exist. Seeds does not create table.")
	} else {
		if err := Db.Migrator().CreateTable(&models.Todo{}); err != nil {
			panic(err)
		}
		var todos models.ToDos
		for i := 1; i <= 10; i++ {
			todos = append(todos, models.Todo{
				Id:      i,
				Name:    fmt.Sprintf("name_%d", i),
				Content: fmt.Sprintf("content_%d", i),
			})
		}
		Db.Create(&todos)
		fmt.Println("create ten todo data!")
	}
}
