package internals

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Task struct {
	Id      int
	Name    string
	Content string
}

func ConnectToDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connection parameters
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB_NAME")

	// Create the connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

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

	for _, task := range tasks {
		fmt.Printf("%d,%s,%s\n", task.Id, task.Name, task.Content)
	}
}
