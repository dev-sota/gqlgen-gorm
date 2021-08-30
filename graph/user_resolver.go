package graph

import (
	"context"
	"fmt"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
)

type userResolver struct{ *Resolver }

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
