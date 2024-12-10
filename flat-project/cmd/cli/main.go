package main

import (
	"taskmanager-elaborated/internal/task"
)

func main() {
	store := task.NewMemoryStore()
	service := task.NewService(store)
	handler := task.NewCLIHandler(service)

	handler.CreateTask()
}
