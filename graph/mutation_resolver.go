package graph

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (string, error) {
	panic(fmt.Errorf("not implemented"))
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
