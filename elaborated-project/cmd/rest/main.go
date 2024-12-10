package main

import (
	"log"
	"net/http"
	"taskmanager-flat/internal/domain"
	"taskmanager-flat/internal/infrastructure/delivery/rest"
	"taskmanager-flat/internal/infrastructure/repository"
)

func main() {
	store := repository.NewMemoryStore()
	service := domain.NewTaskService(store)
	handler := rest.NewRestHandler(service)

	http.HandleFunc("/tasks/create", handler.CreateTask)

	log.Println("REST server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
