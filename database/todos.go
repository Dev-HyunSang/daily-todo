package database

import (
	"context"
	"errors"

	"github.com/dev-hyunsang/daily-todo/ent"
	"github.com/dev-hyunsang/daily-todo/ent/predicate"
	"github.com/dev-hyunsang/daily-todo/ent/todo"
	"github.com/dev-hyunsang/daily-todo/ent/user"
	"github.com/dev-hyunsang/daily-todo/models"
	"github.com/google/uuid"
)

func CreateToDo(todo models.ToDo) (*ent.ToDo, error) {
	clinet, err := ConnectionDB()
	if err != nil {
		return nil, err
	}

	// 입력값 사이즈 검증함. 없는 경우 오류로 반환함.
	if len(todo.UserUUID) == 0 || len(todo.Context) == 0 {
		return nil, errors.New("Failed Not Input Value")
	}

	result, err := clinet.ToDo.Create().
		SetTodoUUID(todo.ToDoUUID).
		SetUserUUID(todo.UserUUID).
		SetIsDone(todo.IsDone).
		SetContext(todo.Context).
		SetCreatedAt(todo.CreatedAt).
		SetUpdatedAt(todo.UpdatedAt).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func EditToDo(userUUID, todoUUID, content string, isDone bool) error {
	parseUserUUID, err := uuid.Parse(userUUID)
	if err != nil {
		return err
	}

	parseToDoUUID, err := uuid.Parse(todoUUID)
	if err != nil {
		return err
	}

	client, err := ConnectionDB()
	if err != nil {
		return err
	}

	if err = client.ToDo.Update().
		Where(todo.UserUUID(parseUserUUID)).
		Where(todo.TodoUUID(parseToDoUUID)).
		SetIsDone(isDone).
		SetContext(content).
		Exec(context.Background()); err != nil {
		return err
	}

	return nil
}

func AllListToDo(userUUID string) ([]*ent.ToDo, error) {
	parseUserUUID, err := uuid.Parse(userUUID)
	if err != nil {
		return nil, err
	}

	client, err := ConnectionDB()
	if err != nil {
		return nil, err
	}

	result, err := client.ToDo.Query().
		Where(todo.UserUUID(parseUserUUID)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteToDo(userUUID, todoUUID string) (int, error) {
	parseUserUUID, err := uuid.Parse(userUUID)
	if err != nil {
		return 0, err
	}

	parseToDoUUID, err := uuid.Parse(todoUUID)
	if err != nil {
		return 0, err
	}

	client, err := ConnectionDB()
	if err != nil {
		return 0, err
	}

	result, err := client.ToDo.Delete().
		Where(predicate.ToDo(user.UserUUID(parseUserUUID))).
		Where(predicate.ToDo(todo.TodoUUID(parseToDoUUID))).
		Exec(context.Background())
	if err != nil {
		return 0, err
	}

	return result, nil
}
