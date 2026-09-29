package main

import (
	"fmt"

	"github.com/jason127vip-dot/SelfTest/client"
)

func main() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTA2MzgxMTcsInVzZXJJZCI6MSwidXNlcm5hbWUiOiJqaW4ifQ.gI2CPQOkNtD_Iz3IpTShL0cMFNcueqDG0JrF2OGeE7M"

	tasks, err := client.GetAllTasks(token)
	if err != nil {
		fmt.Println("调用失败:", err)
		return
	}

	fmt.Println("任务数量:", len(tasks))

	for _, task := range tasks {
		fmt.Println(task.ID, task.Title, task.Status)
	}
}
