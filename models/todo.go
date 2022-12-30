package models

type Todo struct {
	Id      int `gorm:"primary_key"`
	Name    string
	Content string
}

type ToDos []Todo
