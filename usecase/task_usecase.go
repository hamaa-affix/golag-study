package usecase

import "golag-study/model"

type TaskUsecaseInterface interface {
	Index() ([]model.Task, error)
	Show(taskId uint) (model.Task, error)
	Store(task model.Task) (model.Task, error)
	Update(taskId uint, task model.Task) (model.Task, error)
	Delete(taskId uint) error
}
