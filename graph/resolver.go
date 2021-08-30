package graph

import "github.com/dev-sota/gqlgen-gorm/graph/generated"

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
