package models

import (
	"time"

	"github.com/dev-hyunsang/daily-todo/ent"
)

type RequestJoinUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

type RequestLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestCreateToDo struct {
	Context string `json:"context"`
}
type RequestEditToDo struct {
	ToDoUUID string `json:"todo_uuid"`
	IsDone   bool   `json:"is_done"`
	Context  string `json:"context"`
}

type RequestDeleteToDo struct {
	ToDoUUID string `json:"todo_uuid"`
}

// ================================================

type ErrMetaData struct {
	IsSuccess  bool   `json:"is_success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	ErrMessage string `json:"err_message"`
}

type MetaData struct {
	IsSuccess  bool   `json:"is_success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type ResponseDoneJoinUser struct {
	MetaData    `json:"meta_data"`
	Data        *ent.User `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ResponseDoneLoginUser struct {
	MetaData    `json:"meta_data"`
	Data        TokenDetails `json:"data"`
	ResponsedAt time.Time    `json:"responsed_at"`
}

type ResponseDoneLogoutUser struct {
	MetaData    `json:"meta_data"`
	Data        int64     `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ResponseDoneCreateToDo struct {
	MetaData    `json:"meta_data"`
	Data        *ent.ToDo `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ResponseDoneAllListToDo struct {
	MetaData    `json:"meta_data"`
	Data        []*ent.ToDo `json:"data"`
	ResponsedAt time.Time   `json:"responsed_at"`
}

type ResponseDoneDeleteToDo struct {
	MetaData    `json:"meta_data"`
	Data        int       `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ResponseDoneEditToDo struct {
	MetaData    `json:"meta_data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ErrResponse struct {
	ErrMetaData `json:"meta_data"`
	ResponsedAt time.Time `json:"responsed_at"`
}
