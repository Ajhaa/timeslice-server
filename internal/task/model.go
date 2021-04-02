package task

import "gorm.io/gorm"

// Task represents a task
type Task struct {
	gorm.Model
	Name     string
	Duration uint
}
