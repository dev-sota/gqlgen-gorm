package graph

import (
	"context"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type queryResolver struct {
	*Resolver
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var ts []*model.Todo
	if err := r.Resolver.DB.Find(&ts).Error; err != nil {
		return nil, err
	}
	return ts, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	var t model.Todo
	if err := r.Resolver.DB.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var us []*model.User
	if err := r.Resolver.DB.Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var u model.User
	if err := r.Resolver.DB.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
