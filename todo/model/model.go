package model

type TodoItemModel struct {
	Id          int `gorm:"primary_key"`
	Description string
	Completed   string
}