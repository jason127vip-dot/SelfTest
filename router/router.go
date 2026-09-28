package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jason127vip-dot/SelfTest/handler"
	"github.com/jason127vip-dot/SelfTest/middleware"
)

func RegisterRoutes(r *gin.Engine, taskHandler *handler.TaskHandler) {

	r.POST("/login", taskHandler.Login)

	taskGroup := r.Group("/api/tasks")
	taskGroup.Use(middleware.AuthMiddleware())

	taskGroup.GET("", taskHandler.QueryAllTasks)
	taskGroup.POST("/add", taskHandler.CreateTask)
	taskGroup.POST("/update", taskHandler.UpdateTask)
	taskGroup.POST("/delete", taskHandler.DeleteTask)

}
