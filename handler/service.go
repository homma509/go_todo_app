package handler

import (
	"context"

	"github.com/homma509/go_todo_app/entity"
)

//go:generate moq -out moq_test.go . ListTaskService AddTaskService
type ListTaskService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}