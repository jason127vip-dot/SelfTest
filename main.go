package main

import (
	"github.com/jason127vip-dot/SelfTest/model"

	"github.com/gin-gonic/gin"

	"github.com/jason127vip-dot/SelfTest/service"
)

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	r.GET("/task", func(c *gin.Context) {
		tasks := []model.Task{
			{ID: "1", Title: "Task 1", Description: "Description 1", Status: "Pending"},
			{ID: "2", Title: "Task 2", Description: "Description 2", Status: "In Progress"},
			{ID: "3", Title: "Task 3", Description: "Description 3", Status: "Completed"},
		}
		c.JSON(200, tasks)
	})

	r.POST("/saveTask", func(ctx *gin.Context) {
		var task model.Task
		if error := ctx.ShouldBindJSON(&task); error != nil {
			ctx.JSON(400, gin.H{"error": error.Error()})
			return
		}

		ctx.JSON(200, task)
	})

	taskService := &service.TaskServiceImpl{}

	r.GET("/alltasks", func(c *gin.Context) {
		tasks, err := taskService.QueryAllTasks()

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, tasks)
	})

	r.Run(":8080")
}
