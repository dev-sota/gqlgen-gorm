package graph

import (
	"github.com/dev-sota/gqlgen-gorm/graph/generated"
	"gorm.io/gorm"
)

type Resolver struct {
	*gorm.DB
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}
