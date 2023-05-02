package database

import (
	"context"

	"github.com/dev-hyunsang/daily-todo/ent"
	"github.com/dev-hyunsang/daily-todo/ent/user"
	"github.com/dev-hyunsang/daily-todo/models"
)

func CreateUser(data models.User) (*ent.User, error) {
	client, err := ConnectionDB()
	if err != nil {
		return nil, err
	}

	result, err := client.User.Create().
		SetUserUUID(data.UserUUID).
		SetEmail(data.Email).
		SetPassword(data.Password).
		SetNickname(data.NickName).
		SetCreatedAt(data.CreatedAt).
		SetUpdatedAt(data.UpdatedAt).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SearchUserByEmail(email string) (*ent.User, error) {
	client, err := ConnectionDB()
	if err != nil {
		return nil, err
	}

	result, err := client.User.Query().
		Where(user.Email(email)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return result, nil
}
