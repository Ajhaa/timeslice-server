package task

import (
	"context"

	"gorm.io/gorm"
)

// Repository goes brr
type Repository struct {
	DB *gorm.DB
}

func (r Repository) create(ctx context.Context, t partialTask) Task {
	task := Task{Name: t.Name, Duration: t.Duration}
	r.DB.WithContext(ctx).Create(&task)

	return task
}

func (r Repository) findAll(ctx context.Context) []Task {
	var tasks []Task

	r.DB.WithContext(ctx).Find(&tasks)

	return tasks
}
