package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"

	"github.com/mishra811/go-htmx-app/internals"
)

type Task struct {
	Id      int
	Name    string
	Content string
}

var db = internals.ConnectToDB()

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/create", createTask)
	http.HandleFunc("/delete/", deleteTaskById)
	http.HandleFunc("/delete-all", clearAllTasks)
	http.HandleFunc("/show-form/", showUpdateForm)
	http.HandleFunc("/update-task/", updateTask)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("templates/index.html"))
	tasks := getAllTasks()

	templ.Execute(w, tasks)
}

func getAllTasks() []*Task {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", "tasks"))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	tasks := make([]*Task, 0)

	for rows.Next() {
		task := new(Task)
		err := rows.Scan(&task.Id, &task.Name, &task.Content)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tasks
}

func getTaskById(taskId string) *Task {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE Id = %s", "tasks", taskId))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	Tasks := make([]*Task, 0)

	for rows.Next() {
		task := new(Task)
		err := rows.Scan(&task.Id, &task.Name, &task.Content)
		if err != nil {
			log.Fatal(err)
		}

		Tasks = append(Tasks, task)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return Tasks[0]
}

func createTask(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	content := r.PostFormValue("content")
	templ := template.Must(template.ParseFiles("templates/index.html"))

	_, err := db.Query(
		fmt.Sprintf("INSERT INTO tasks (Name,Content) VALUES ('%s','%s') ;", name, content),
	)
	if err != nil {
		log.Fatal(err, http.StatusInternalServerError)
		return
	}

	templ.ExecuteTemplate(w,
		"film-list-element",
		Task{Id: 0, Name: name, Content: content},
	)
}

func deleteTaskById(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("templates/index.html"))
	taskId := strings.Split(r.URL.Path, "/")[2]

	_, err := db.Query(fmt.Sprintf("DELETE FROM tasks WHERE Id = %s", taskId))
	if err != nil {
		log.Fatal(err, http.StatusInternalServerError)
		return
	}

	Tasks := getAllTasks()

	templ.ExecuteTemplate(w, "task-list", Tasks)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("templates/index.html"))
	taskId := strings.Split(r.URL.Path, "/")[2]

	name := r.PostFormValue("name")
	content := r.PostFormValue("content")
	_, err := db.Query(
		fmt.Sprintf(
			"UPDATE tasks SET Name = '%s', Content = '%s' WHERE Id = %s ",
			name,
			content,
			taskId,
		),
	)
	if err != nil {
		log.Fatal(err, http.StatusInternalServerError)
		return
	}
	Tasks := getAllTasks()

	templ.ExecuteTemplate(w, "task-list", Tasks)
}

func showUpdateForm(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("templates/update-form.html"))
	taskId := strings.Split(r.URL.Path, "/")[2]
	Task := getTaskById(taskId)

	templ.ExecuteTemplate(w, "film-update-element", Task)
}

func clearAllTasks(w http.ResponseWriter, r *http.Request) {
	_, err := db.Query(fmt.Sprintf("DELETE FROM tasks ;"))
	if err != nil {
		log.Fatal(err, http.StatusInternalServerError)
		return
	}

	Tasks := getAllTasks()
	templ := template.Must(template.ParseFiles("templates/index.html"))

	templ.ExecuteTemplate(w, "film-list", Tasks)
}
