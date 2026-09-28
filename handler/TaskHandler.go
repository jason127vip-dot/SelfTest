package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jason127vip-dot/SelfTest/dto"
	"github.com/jason127vip-dot/SelfTest/model"
	"github.com/jason127vip-dot/SelfTest/response"
	"github.com/jason127vip-dot/SelfTest/service"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	var req dto.CreateTaskRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, err.Error())
		return
	}

	task := model.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	err := h.service.SaveTask(&task)
	if err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}

	response.Success(ctx, task)
}

func (h *TaskHandler) UpdateTask(ctx *gin.Context) {
	var req dto.UpdateTaskRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, err.Error())
		return
	}

	task := model.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	updatedTask, err := h.service.UpdateTask(&task)
	if err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, updatedTask)
}

func (h *TaskHandler) DeleteTask(ctx *gin.Context) {
	var req dto.DeleteTaskRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, err.Error())
		return
	}

	task := model.Task{
		ID: req.ID,
	}

	err := h.service.DeleteTask(&task)
	if err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}

	response.Success(ctx, gin.H{"message": "Task deleted successfully"})
}

func (h *TaskHandler) QueryAllTasks(c *gin.Context) {
	status := c.Query("status")

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	tasks, err := h.service.QueryAllTasks(status, page, pageSize)

	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, tasks)
}
