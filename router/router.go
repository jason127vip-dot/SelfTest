package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jason127vip-dot/SelfTest/handler"
)

func RegisterRoutes(r *gin.Engine, taskHandler *handler.TaskHandler) {

	taskGroup := r.Group("/api/tasks")

	taskGroup.GET("", taskHandler.QueryAllTasks)
	taskGroup.POST("/add", taskHandler.CreateTask)
	taskGroup.POST("/update", taskHandler.UpdateTask)
	taskGroup.POST("/delete", taskHandler.DeleteTask)

}
