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

// Session is the resolver for the session field.
func (r *userResolver) Session(ctx context.Context, obj *ent.User) ([]*ent.Session, error) {
	panic(fmt.Errorf("not implemented: Session - session"))
}

// PasswordResetRequest is the resolver for the passwordResetRequest field.
func (r *userResolver) PasswordResetRequest(ctx context.Context, obj *ent.User) ([]*ent.PasswordResetRequest, error) {
	panic(fmt.Errorf("not implemented: PasswordResetRequest - passwordResetRequest"))
}

// Credentials is the resolver for the credentials field.
func (r *userResolver) Credentials(ctx context.Context, obj *ent.User) ([]*ent.Credentials, error) {
	panic(fmt.Errorf("not implemented: Credentials - credentials"))
}

// Token is the resolver for the token field.
func (r *userResolver) Token(ctx context.Context, obj *ent.User) ([]*ent.Token, error) {
	panic(fmt.Errorf("not implemented: Token - token"))
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
