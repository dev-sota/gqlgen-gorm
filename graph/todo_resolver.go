package graph

import (
	"context"

	"github.com/dev-sota/gqlgen-gorm/dataloader/graph"
	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj == nil {
		return nil, nil
	}
	return graph.For(ctx).UserById.Load(obj.UserID)
}
