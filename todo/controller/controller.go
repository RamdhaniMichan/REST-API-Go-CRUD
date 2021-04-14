package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"

	"todolist-app/todo/config"
	"todolist-app/todo/model"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Add new TodoItem. Saving to database.")
	todo := &model.TodoItemModel{Description: description, Completed: "incompleted"}
	config.Connect().Create(&todo)
	result := config.Connect().Last(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)
	
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := GetItemByID(id)
	if err == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"updated": false, "error": "Record Not Found"}`)
	} else {
		completed := r.FormValue("completed")
		log.WithFields(log.Fields{"Id": id, "Completed": completed}).Info("Updating Todo Item")
		todo := &model.TodoItemModel{}
		config.Connect().First(&todo, id)
		todo.Completed = completed
		config.Connect().Save(&todo)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"updated": true}`)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := GetItemByID(id)
	if err == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": false, "error": "Record Not Found"}`)
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Deleting Todo Item")
		todo := &model.TodoItemModel{}
		config.Connect().First(&todo, id)
		config.Connect().Delete(&todo)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": true}`)
	}
}

func GetItemByID(Id int) bool {
	todo := &model.TodoItemModel{}
	result := config.Connect().First(&todo, Id)
	if result.Error != nil {
		log.Warn("Todo Item not Found")
	} 
	return true
}

func GetCompletedItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Get Completed Todo Items")
	completedTodoItems := GetTodoItems("completed")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completedTodoItems)
}

func GetProcessItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Get Process Todo Items")
	processTodoItems := GetTodoItems("process")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(processTodoItems)
}

func GetIncompletedItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Get Incompleted Todo Items")
	IncompletedTodoItems := GetTodoItems("incompleted")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(IncompletedTodoItems)
}

func GetTodoItems(completed string) interface{} {
	var todos []model.TodoItemModel
	TodoItems := config.Connect().Where("completed = ?", completed).Find(&todos).Value
	return TodoItems
}

func Health(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive":true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}