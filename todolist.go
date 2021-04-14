package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"

	"todolist-app/todo/config"
	"todolist-app/todo/controller"
	"todolist-app/todo/model"
)

func main() {
	defer config.Connect().Close()

	config.Connect().Debug().DropTableIfExists(&model.TodoItemModel{})
	config.Connect().Debug().AutoMigrate(&model.TodoItemModel{})

	log.Info("Starting Todo List API Server")
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Health).Methods("GET")
	router.HandleFunc("/todo-completed", controller.GetCompletedItems).Methods("GET")
	router.HandleFunc("/todo-incomplete", controller.GetIncompletedItems).Methods("GET")
	router.HandleFunc("/todo-process", controller.GetProcessItems).Methods("GET")
	router.HandleFunc("/todo", controller.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", controller.UpdateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", controller.DeleteItem).Methods("DELETE")

	http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"})) (router))
}