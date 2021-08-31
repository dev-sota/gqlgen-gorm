package graph

import (
	"context"

	"github.com/dev-sota/gqlgen-gorm/dataloader/graph"
	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type userResolver struct {
	*Resolver
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	if obj == nil {
		return []*model.Todo{}, nil
	}
	return graph.For(ctx).TodosByUserIDs.Load(obj.ID)
}
