package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"jet/ent"
	graphql1 "jet/graphql"
)

// ID is the resolver for the id field.
func (r *accountResolver) ID(ctx context.Context, obj *ent.Account) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Role is the resolver for the role field.
func (r *accountResolver) Role(ctx context.Context, obj *ent.Account) (ent.Role, error) {
	panic(fmt.Errorf("not implemented: Role - role"))
}

// Account returns graphql1.AccountResolver implementation.
func (r *Resolver) Account() graphql1.AccountResolver { return &accountResolver{r} }

type accountResolver struct{ *Resolver }
