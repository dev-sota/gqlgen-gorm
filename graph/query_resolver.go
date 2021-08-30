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
	panic(fmt.Errorf("not implemented"))
}
