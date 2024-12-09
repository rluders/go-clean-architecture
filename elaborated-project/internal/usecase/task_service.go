package usecase

import "taskmanager/internal/domain"

type TaskService struct {
	repo domain.TaskRepositorier
}

func NewTaskService(repo domain.TaskRepositorier) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task domain.Task) error {
	if err := task.Validate(); err != nil {
		return err
	}
	return s.repo.Save(task)
}

func (s *TaskService) GetTasks() ([]domain.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) CompleteTask(id int) error {
	return s.repo.MarkAsComplete(id)
}
