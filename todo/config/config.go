package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB{
	var err error

	db, err := gorm.Open("mysql", "root:@/todo-list?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	fmt.Println("Database Conect")
	
	return db

}