package service

import (
	"context"
	"fmt"

	"github.com/ShinnosukeSuzuki/web-app-develop-golang-todo/entity"
	"github.com/jmoiron/sqlx"
)

type ListTask struct {
	DB   *sqlx.DB
	Repo TaskLister
}

func (lt *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return tasks, nil
}
