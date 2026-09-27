package service

import (
	"github.com/jason127vip-dot/SelfTest/model"
)

type TaskService interface {
	QueryAllTasks() ([]model.Task, error)
}
