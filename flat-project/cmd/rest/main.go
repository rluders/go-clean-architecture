package main

import (
	"log"
	"net/http"
	"taskmanager-elaborated/internal/task"
)

func main() {
	store := task.NewMemoryStore()
	service := task.NewService(store)
	handler := task.NewRestHandler(service)

	http.HandleFunc("/tasks/create", handler.CreateTask)

	log.Println("REST server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
