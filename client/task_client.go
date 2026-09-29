package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jason127vip-dot/SelfTest/model"
)

type TaskResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []model.Task `json:"data"`
}

func GetAllTasks(token string) ([]model.Task, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/api/tasks",
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", string(body))
	}

	var result TaskResponse

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
