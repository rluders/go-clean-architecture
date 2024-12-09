package main

import (
	"log"
	"net/http"
	"taskmanager/internal/infrastructure/delivery/rest"
	"taskmanager/internal/infrastructure/repository"
	"taskmanager/internal/usecase"
)

func main() {
	store := repository.NewMemoryStore()
	service := usecase.NewTaskService(store)
	handler := rest.NewRestHandler(service)

	http.HandleFunc("/tasks/create", handler.CreateTask)

	log.Println("REST server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
