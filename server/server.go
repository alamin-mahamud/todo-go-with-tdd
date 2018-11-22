package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alamin-mahamud/todo-go-with-tdd/todo"
	"github.com/gorilla/mux"
)

var taskManager = todo.NewTaskManager()

func listTasks(w http.ResponseWriter, r *http.Request) {
	res := struct{ tasks []*todo.Task }{}
	res.tasks = taskManager.All()
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Oops!", http.StatusInternalServerError)
		log.Println(err)
	}
}

func newTask(w http.ResponseWriter, r *http.Request) {
	res := struct{ Title string }{}
	res.Title
}

func registerRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/task", listTasks).Methods("GET")
	r.HandleFunc("/task", newTask).Methods("POST")
	r.HandleFunc("/task/{id}", getTask).Methods("GET")
	r.HandleFunc("/task/{id}", updateTasks).Methods("PUT")
	return r
}

func main() {
	r := mux.NewRouter()
	r = registerRoutes(r)
	http.ListenAndServe(":8080", r)
}
