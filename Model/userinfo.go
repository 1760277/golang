package Model

import (
	"./origin"
)
type UserInfo struct {
	UserID string `json:"user_id" form:"user_id"`
	UserName string `json:"user_name" form:"user_name"`
	UserGroup string `json:"user_group form:"user_name"`
	UserPassword string `json:"user_password" form:"user_password"`
}

type Users struct {
	ListUser = []UserInfo
}

func getUsers()