package domain

type TaskRepositorier interface {
	GetAll() ([]Task, error)
	Save(task Task) error
	MarkAsComplete(id int) error
}
