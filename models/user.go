package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserUUID  uuid.UUID `json:"user_uuid"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	NickName  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TokenDetails struct {
	UserUUID     uuid.UUID `json:"user_uuid"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	AccessUUID   uuid.UUID `json:"access_uuid"`
	RefreshUUID  uuid.UUID `json:"refresh_uuid"`
	AtExpires    int64     `json:"at_expires"`
	RtExpires    int64     `json:"rt_expires"`
}

type AccessDetails struct {
	AccessUUID string `json:"access_uuid"`
	UserUUID   string `json:"user_uuid"`
}
