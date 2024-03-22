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
func (r *passwordResetRequestResolver) User(ctx context.Context, obj *ent.PasswordResetRequest) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// PasswordResetRequest returns graph.PasswordResetRequestResolver implementation.
func (r *Resolver) PasswordResetRequest() graph.PasswordResetRequestResolver {
	return &passwordResetRequestResolver{r}
}

type passwordResetRequestResolver struct{ *Resolver }
