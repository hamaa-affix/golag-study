package repository

import (
	"golag-study/model"

	"gorm.io/gorm"
)

type TaskRepositoryInterface interface {
	FetchAll(tasks *[]model.Task) error
	FindById(taskId uint, task *model.Task) error
	Create(task *model.Task) error
	Update(taskId uint, task *model.Task) error
	Delete(taskId uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func (t *TaskRepository) FetchAll(tasks *[]model.Task) error {
	return nil
}

func (t *TaskRepository) FindById(taskId uint, task *model.Task) error {
	return nil
}

func (t *TaskRepository) Create(task *model.Task) error {
	return nil
}

func (t *TaskRepository) Update(taskId uint, task *model.Task) error {
	return nil
}

func (t *TaskRepository) Delete(taskId uint) error {
	return nil
}
