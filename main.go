package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mishra811/go-htmx-app/internals"
)

type Task internals.Task

func main() {
	internals.ConnectToDB()
	http.HandleFunc("/", handler)
	http.HandleFunc("/create", createTask)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))

	tasks := []Task{
		{Id: 1, Name: "cooking", Content: "cook breakfast"},
	}

	templ.Execute(w, tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	content := r.PostFormValue("content")
	templ := template.Must(template.ParseFiles("index.html"))

	templ.ExecuteTemplate(w,
		"film-list-element",
		Task{Id: 2, Name: name, Content: content},
	)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
}

func deleteTaskById(w http.ResponseWriter, r *http.Request) {
}

func clearAllTasks(w http.ResponseWriter, r *http.Request) {
}
