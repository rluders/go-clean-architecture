package main

import (
	"taskmanager/internal/infrastructure/delivery/cli"
	"taskmanager/internal/infrastructure/repository"
	"taskmanager/internal/usecase"
)

func main() {
	store := repository.NewMemoryStore()
	service := usecase.NewTaskService(store)
	handler := cli.NewCLIHandler(service)

	handler.CreateTask()
}
