package service

import (
	"github.com/jason127vip-dot/goself/model"
)

type TaskService interface {
	QueryAllTasks() ([]model.Task, error)
}
