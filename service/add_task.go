package service

import (
	"context"
	"fmt"

	"github.com/ShinnosukeSuzuki/web-app-develop-golang-todo/entity"
	"github.com/jmoiron/sqlx"
)

type AddTask struct {
	DB   *sqlx.DB
	Repo TaskAdder
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	t := &entity.Task{
		Title:  title,
		Status: entity.TaskStatusToDo,
	}

	err := a.Repo.AddTask(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
