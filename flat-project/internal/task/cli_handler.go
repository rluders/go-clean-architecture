package task

import (
	"bufio"
	"fmt"
	"os"
)

type CLIHandler struct {
	service *Service
}

func NewCLIHandler(service *Service) *CLIHandler {
	return &CLIHandler{service: service}
}

func (h *CLIHandler) CreateTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')

	task := Task{Title: title}
	if err := h.service.CreateTask(task); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Task created successfully!")
	}
}
