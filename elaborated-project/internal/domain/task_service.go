package domain

type TaskService struct {
	repo TaskRepositorier
}

func NewTaskService(repo TaskRepositorier) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) error {
	if err := task.Validate(); err != nil {
		return err
	}
	return s.repo.Save(task)
}

func (s *TaskService) GetTasks() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) CompleteTask(id int) error {
	return s.repo.MarkAsComplete(id)
}
