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

// ================================================

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

type ErrResponse struct {
	MetaData    `json:"meta_data"`
	ResponsedAt time.Time `json:"responsed_at"`
}
