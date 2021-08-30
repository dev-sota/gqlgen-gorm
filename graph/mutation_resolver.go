package graph

import (
	"context"
	"strconv"
	"time"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (string, error) {
	var m model.Todo
	m.ID = timeStampID()
	m.Text = input.Text
	m.UserID = input.UserID

	if err := r.Resolver.DB.Create(&m).Error; err != nil {
		return "", err
	}
	return m.ID, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var u model.User
	u.ID = timeStampID()
	u.Name = input.Name

	if err := r.Resolver.DB.Create(&u).Error; err != nil {
		return "", err
	}
	return u.ID, nil
}

func timeStampID() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
