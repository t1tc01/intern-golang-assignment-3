package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"gitlab.com/hedwig-phan/assignment-3/ent"
	"gitlab.com/hedwig-phan/assignment-3/graph"
)

// User is the resolver for the user field.
func (r *tokenResolver) User(ctx context.Context, obj *ent.Token) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Token returns graph.TokenResolver implementation.
func (r *Resolver) Token() graph.TokenResolver { return &tokenResolver{r} }

type tokenResolver struct{ *Resolver }
