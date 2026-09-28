package service

import (
	"github.com/jason127vip-dot/SelfTest/model"
)

type TaskService interface {
	QueryAllTasks(status string,
		page int,
		pageSize int) ([]model.Task, error)

	SaveTask(task *model.Task) error

	UpdateTask(task *model.Task) (*model.Task, error)

	DeleteTask(task *model.Task) error
}
