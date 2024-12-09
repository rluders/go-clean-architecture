package task

type Repositorier interface {
	GetAll() ([]Task, error)
	Save(task Task) error
	MarkAsComplete(id int) error
}
