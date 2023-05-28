package repository

import (
	"golag-study/model"

	"gorm.io/gorm"
)
type taskRepository interface {
	FetchALl() ([]model.Task, error)
}

type TaksRepo struct {
	db *gorm.DB
}

func (t *TaksRepo) FetchALl() ([]model.Task, error) {
	return []model.Task{}, nil
}
