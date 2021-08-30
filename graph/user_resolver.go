package graph

import (
	"context"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type userResolver struct {
	*Resolver
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	if obj == nil {
		return []*model.Todo{}, nil
	}
	var ts []*model.Todo

	err := r.Resolver.DB.Where("user_id = ?", obj.ID).Find(&ts).Error
	if err != nil {
		return nil, err
	}

	return ts, nil
}
