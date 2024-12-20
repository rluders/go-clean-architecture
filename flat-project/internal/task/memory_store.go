package task

import (
	"errors"
	"sync"
)

type MemoryStore struct {
	tasks []Task
	mu    sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{tasks: []Task{}}
}

func (m *MemoryStore) GetAll() ([]Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.tasks, nil
}

func (m *MemoryStore) Save(task Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	task.ID = len(m.tasks) + 1
	m.tasks = append(m.tasks, task)
	return nil
}

func (m *MemoryStore) MarkAsComplete(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks[i].Completed = true
			return nil
		}
	}
	return errors.New("task not found")
}
