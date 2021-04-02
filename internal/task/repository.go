package task

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(ctx context.Context) []Task
	Create(ctx context.Context, input PartialTask) Task
}

// Repository goes brr
type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{DB: db}
}

func (r repository) Create(ctx context.Context, t PartialTask) Task {
	task := Task{Name: t.Name, Duration: t.Duration}
	r.DB.WithContext(ctx).Create(&task)

	return task
}

func (r repository) FindAll(ctx context.Context) []Task {
	var tasks []Task

	r.DB.WithContext(ctx).Find(&tasks)

	return tasks
}
