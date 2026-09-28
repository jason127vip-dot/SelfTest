package dto

type DeleteTaskRequest struct {
	ID uint `json:"id" binding:"required"`
}
