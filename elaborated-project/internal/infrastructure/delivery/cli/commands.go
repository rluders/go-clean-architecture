package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"taskmanager/internal/domain"
	"taskmanager/internal/usecase"
	"time"
)

type CLIHandler struct {
	service *usecase.TaskService
}

func NewCLIHandler(service *usecase.TaskService) *CLIHandler {
	return &CLIHandler{service: service}
}

func (h *CLIHandler) CreateTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')

	task := domain.Task{Title: title, DueDate: time.Now().AddDate(0, 0, 7)}
	if err := h.service.CreateTask(task); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Task created successfully!")
		output, _ := json.Marshal(task)
		fmt.Printf("Result: %s", output)
	}
}
