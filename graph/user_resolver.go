package graph

import (
	"context"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type userResolver struct{ *Resolver }

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "1111",
			Text: "example",
			Done: true,
		},
	}, nil
}
