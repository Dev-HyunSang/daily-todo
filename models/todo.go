package models

import (
	"time"

	"github.com/google/uuid"
)

type ToDo struct {
	ToDoUUID  uuid.UUID `json:"todo_uuid"`
	UserUUID  uuid.UUID `json:"user_uuid"`
	IsDone    bool      `json:"is_done"`
	Context   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"edited_at"`
}
