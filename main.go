package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Task struct {
	Id          int
	Content     string
	isCompleted bool
	createAt    time.Duration
	updatedAt   time.Duration
}

type tasks []Task

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/create", createTask)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("./index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		"Welcome to my page",
		"This is a todo app",
	}

	err = templ.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
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
