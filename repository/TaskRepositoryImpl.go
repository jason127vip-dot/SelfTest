package repository

import (
	"github.com/jason127vip-dot/SelfTest/model"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{
		db: db,
	}
}

func (r *TaskRepositoryImpl) QueryAllTasks(status string,
	page int,
	pageSize int) ([]model.Task, error) {
	var tasks []model.Task

	query := r.db
	if status != "" {
		query = query.Where("status = ?", status)
	}

	offset := (page - 1) * pageSize

	result := query.
		Limit(pageSize).
		Offset(offset).
		Order("id asc").
		Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (r *TaskRepositoryImpl) SaveTask(task *model.Task) error {
	result := r.db.Create(task)

	if result.Error != nil {
		return result.Error
	}

	return result.Error
}

func (r *TaskRepositoryImpl) UpdateTask(newTask *model.Task) (*model.Task, error) {
	var oldTask model.Task

	// 根据 newTask.ID 查询数据库原数据
	if err := r.db.First(&oldTask, "id = ?", newTask.ID).Error; err != nil {
		return nil, err
	}
	// 更新字段
	oldTask.Title = newTask.Title
	oldTask.Description = newTask.Description
	oldTask.Status = newTask.Status
	// 保存
	if err := r.db.Save(&oldTask).Error; err != nil {
		return nil, err
	}
	return &oldTask, nil

}
func (r *TaskRepositoryImpl) DeleteTask(task *model.Task) error {
	result := r.db.Delete(&model.Task{}, "id = ?", task.ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
