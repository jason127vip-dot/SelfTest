package service

import "github.com/jason127vip-dot/SelfTest/model"

type TaskServiceImpl struct {
}

func (s *TaskServiceImpl) QueryAllTasks() ([]model.Task, error) {
	tasks := []model.Task{
		{
			ID:          "1",
			Title:       "Learn Go",
			Description: "Study Gin",
			Status:      "todo",
		},
		{
			ID:          "2",
			Title:       "Build API",
			Description: "Practice service layer",
			Status:      "doing",
		},
	}

	return tasks, nil
}
