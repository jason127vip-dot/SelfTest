package service

import (
	"context"

	"github.com/jason127vip-dot/SelfTest/model"
)

type TaskService interface {
	QueryAllTasks(ctx context.Context, status string,
		page int,
		pageSize int) ([]model.Task, error)

	SaveTask(task *model.Task) error

	UpdateTask(task *model.Task) (*model.Task, error)

	DeleteTask(task *model.Task) error
}
