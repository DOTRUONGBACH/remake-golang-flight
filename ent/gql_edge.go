// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Account) AccOwner(ctx context.Context) (*Customer, error) {
	result, err := a.Edges.AccOwnerOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryAccOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Customer) Accounts(ctx context.Context) (result []*Account, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedAccounts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.AccountsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryAccounts().All(ctx)
	}
	return result, err
}