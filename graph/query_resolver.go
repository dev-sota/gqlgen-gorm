package graph

import (
	"context"
	"fmt"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type queryResolver struct {
	*Resolver
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "001",
			Text: "cleaning",
			Done: false,
			User: &model.User{
				ID:   "user001",
				Name: "Satoshi",
			},
		},
		{
			ID:   "002",
			Text: "shopping",
			Done: true,
			User: &model.User{
				ID:   "user002",
				Name: "Ota",
			},
		},
	}, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
