package repository

import "github.com/jason127vip-dot/SelfTest/model"

type TaskRepository interface {
	QueryAllTasks(status string, page int, pageSize int) ([]model.Task, error)

	SaveTask(task *model.Task) error

	UpdateTask(newTask *model.Task) (*model.Task, error)

	DeleteTask(task *model.Task) error
}
