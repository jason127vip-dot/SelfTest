package main

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/jason127vip-dot/SelfTest/config"
	"github.com/jason127vip-dot/SelfTest/handler"
	"github.com/jason127vip-dot/SelfTest/logger"
	"github.com/jason127vip-dot/SelfTest/middleware"
	"github.com/jason127vip-dot/SelfTest/repository"
	"github.com/jason127vip-dot/SelfTest/router"

	"github.com/gin-gonic/gin"

	"github.com/jason127vip-dot/SelfTest/service"
)

func main() {

	logger.InitLogger()

	slog.Info("server starting")

	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())

	cfg := config.LoadConfig()

	db, err := config.InitDB(cfg.DatabaseURL)

	if err != nil {
		slog.Error(
			"database connect failed",
			"error", err,
		)
		panic(err)
	}

	slog.Info(
		"database connected",
		"connected", db != nil,
	)

	//err = db.AutoMigrate(&model.Task{})
	//if err != nil {
	//panic(err)
	//}

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	router.RegisterRoutes(r, taskHandler)

	/**go printTask()

	fmt.Println("main continues")

	time.Sleep(6 * time.Second)**/

	// var wg sync.WaitGroup

	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1)

	// 	go work(i, &wg)
	// }

	// wg.Wait()

	// fmt.Println("all finished")

}

func printTask() {
	for i := 1; i <= 2; i++ {
		fmt.Println("task:", i)
		time.Sleep(time.Second)
	}
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("start:", id)

	time.Sleep(time.Second)

	fmt.Println("end:", id)
}
