package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jason127vip-dot/SelfTest/dto"
	"github.com/jason127vip-dot/SelfTest/model"
	"github.com/jason127vip-dot/SelfTest/response"
	"github.com/jason127vip-dot/SelfTest/service"
	"github.com/jason127vip-dot/SelfTest/utils"
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

	userID, _ := c.Get("userId")
	username, _ := c.Get("username")

	fmt.Println(userID)
	fmt.Println(username)

	status := c.Query("status")

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		3*time.Second,
	)
	defer cancel()

	tasks, err := h.service.QueryAllTasks(ctx, status, page, pageSize)

	if errors.Is(err, context.DeadlineExceeded) {
		response.Error(c, 504, "request timeout")
		return
	}

	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, tasks)
}

func (h *TaskHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	if req.Username != "jin" || req.Password != "123456" {
		response.Error(c, 401, "invalid username or password")
		return
	}

	token, err := utils.GenerateToken(1, req.Username)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{
		"token": token,
	})
}
