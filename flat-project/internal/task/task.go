package task

import (
	"errors"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    int
	DueDate     time.Time
	Completed   bool
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	if t.DueDate.Before(time.Now()) {
		return errors.New("due date cannot be in the past")
	}
	return nil
}
