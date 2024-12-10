package main

import (
	"taskmanager-flat/internal/domain"
	"taskmanager-flat/internal/infrastructure/delivery/cli"
	"taskmanager-flat/internal/infrastructure/repository"
)

func main() {
	store := repository.NewMemoryStore()
	service := domain.NewTaskService(store)
	handler := cli.NewCLIHandler(service)

	handler.CreateTask()
}
