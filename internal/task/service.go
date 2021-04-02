package task

import (
	"context"

	"gorm.io/gorm"
)

type Service interface {
	Query(ctx context.Context) []Task
	Create(ctx context.Context, input PartialTask) Task
}

type service struct {
	repo Repository
}

type PartialTask struct {
	Name     string `json:"name"`
	Duration uint   `json:"duration"`
}

// NewService creates a new task service
func NewService(db *gorm.DB) Service {
	return service{repo: NewRepository(db)}
}

func (s service) Query(ctx context.Context) []Task {
	return s.repo.FindAll(ctx)
}

func (s service) Create(ctx context.Context, input PartialTask) Task {
	return s.repo.Create(ctx, input)
}
