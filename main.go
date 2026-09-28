package main

import (
	"github.com/jason127vip-dot/SelfTest/config"
	"github.com/jason127vip-dot/SelfTest/handler"
	"github.com/jason127vip-dot/SelfTest/model"
	"github.com/jason127vip-dot/SelfTest/repository"

	"github.com/gin-gonic/gin"

	"github.com/jason127vip-dot/SelfTest/service"
)

func main() {
	r := gin.Default()
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Task{})
	if err != nil {
		panic(err)
	}

	taskRepository := repository.NewTaskRepository(db)

	taskService := service.NewTaskService(taskRepository)

	taskHandler := handler.NewTaskHandler(taskService)

	r.POST("/saveTask", taskHandler.CreateTask)

	r.POST("/deleteTask", taskHandler.DeleteTask)

	r.POST("/updateTask", taskHandler.UpdateTask)

	r.GET("/alltasks", taskHandler.QueryAllTasks)

	r.Run(":8080")
}
