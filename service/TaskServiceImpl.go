package service

import (
	"context"

	"github.com/jason127vip-dot/SelfTest/model"
	"github.com/jason127vip-dot/SelfTest/repository"
)

type TaskServiceImpl struct {
	repository repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		repository: repo,
	}
}

func (s *TaskServiceImpl) QueryAllTasks(ctx context.Context, status string,
	page int,
	pageSize int) ([]model.Task, error) {

	return s.repository.QueryAllTasks(ctx, status, page, pageSize)
}

func (s *TaskServiceImpl) SaveTask(task *model.Task) error {
	return s.repository.SaveTask(task)
}

func (s *TaskServiceImpl) UpdateTask(task *model.Task) (*model.Task, error) {
	return s.repository.UpdateTask(task)
}

func (s *TaskServiceImpl) DeleteTask(task *model.Task) error {
	return s.repository.DeleteTask(task)
}
