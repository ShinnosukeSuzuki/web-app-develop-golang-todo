package service

import (
	"context"

	"github.com/ShinnosukeSuzuki/web-app-develop-golang-todo/entity"
	"github.com/ShinnosukeSuzuki/web-app-develop-golang-todo/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister UserRegisterer
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}

type UserRegisterer interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}
