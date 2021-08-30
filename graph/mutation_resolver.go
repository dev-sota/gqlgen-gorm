package graph

import (
	"context"
	"fmt"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type mutationResolver struct {
	*Resolver
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
