package graph

import (
	"context"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj == nil {
		return nil, nil
	}

	var u model.User
	err := r.Resolver.DB.First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
