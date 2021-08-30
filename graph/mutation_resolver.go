package graph

import (
	"context"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type mutationResolver struct {
	*Resolver
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
		ID:   "001",
		Text: "cleaning",
		Done: false,
		User: &model.User{
			ID:   "user001",
			Name: "Satoshi",
		},
	}, nil
}
