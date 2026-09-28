package main

import (
	"github.com/jason127vip-dot/SelfTest/config"
	"github.com/jason127vip-dot/SelfTest/handler"
	"github.com/jason127vip-dot/SelfTest/middleware"
	"github.com/jason127vip-dot/SelfTest/repository"
	"github.com/jason127vip-dot/SelfTest/router"

	"github.com/gin-gonic/gin"

	"github.com/jason127vip-dot/SelfTest/service"
)

func main() {
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())

	cfg := config.LoadConfig()

	db, err := config.InitDB(cfg.DatabaseURL)

	if err != nil {
		panic(err)
	}
	//err = db.AutoMigrate(&model.Task{})
	//if err != nil {
	//panic(err)
	//}

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	router.RegisterRoutes(r, taskHandler)

	r.Run(":" + cfg.ServerPort)
}
