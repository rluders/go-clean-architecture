package repository

import (
	"errors"
	"sync"
	"taskmanager-flat/internal/domain"
)

type MemoryStore struct {
	tasks []domain.Task
	mu    sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{tasks: []domain.Task{}}
}

func (m *MemoryStore) GetAll() ([]domain.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.tasks, nil
}

func (m *MemoryStore) Save(task domain.Task) error {
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
