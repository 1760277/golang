package Model

import (
	"fmt"

	"./Server"
)

type UserInfo struct {
	UserID       string `json:"user_id" form:"user_id"`
	UserName     string `json:"user_name" form:"user_name"`
	UserGroup    string `json:"user_group form:"user_name"`
	UserPassword string `json:"user_password" form:"user_password"`
}

type Users struct {
	ListUser []UserInfo
}

func getUsers(data Server.DbConfig) error {
	u.data.ConnectDB()
	row, err := u.data.Query(`
		SELECT user_id, user_username, user_group, user_password
		FROM tbl_users
	`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := repositorySummary{}
		err = rows.Scan(
			&repo.UserID,
			&repo.UserName,
			&repo.UserGroup,
			&repo.UserPassword,
		)
		if err != nil {
			return err
		}
		repos.Repositories = append(repos.Repositories, repo)
	}
	fmt.Println(repos.Repositories)
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
