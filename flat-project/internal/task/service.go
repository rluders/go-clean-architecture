package task

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(task Task) error {
	if err := task.Validate(); err != nil {
		return err
	}
	return s.repo.Save(task)
}

func (s *Service) GetTasks() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *Service) CompleteTask(id int) error {
	return s.repo.MarkAsComplete(id)
}
