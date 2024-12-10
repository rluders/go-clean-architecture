package task

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

	task := Task{Title: title, DueDate: time.Now().AddDate(0, 0, 7)}
	if err := h.service.CreateTask(task); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Task created successfully!")
	}
}
