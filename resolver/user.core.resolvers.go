package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"gitlab.com/hedwig-phan/assignment-3/cmd/db"
	"gitlab.com/hedwig-phan/assignment-3/cmd/jwt"
	"gitlab.com/hedwig-phan/assignment-3/config"
	"gitlab.com/hedwig-phan/assignment-3/ent"
	"gitlab.com/hedwig-phan/assignment-3/ent/user"
	"gitlab.com/hedwig-phan/assignment-3/graph"
	"gitlab.com/hedwig-phan/assignment-3/internal/users"
	"gitlab.com/hedwig-phan/assignment-3/internal/util"
	"gitlab.com/hedwig-phan/assignment-3/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (string, error) {

	username := input.Username
	password := input.Password
	email := input.Email

	if username == "" || password == "" {
		return "", &users.InvalidUsernameOrPasswordError{}
	}

	if email == "" {
		return "", &users.InvalidEmailError{}
	}

	//
	user, err := users.CreateUser(ctx, username, email, password, config.ROLE_USER)

	if err != nil {
		return "", err
	}

	//Token
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	username := input.Username
	password := input.Password

	correct := users.AuthenticateLogin(ctx, username, password)
	if !correct {
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	user := db.Client.User.Query().Where(user.Username(username)).OnlyX(ctx)

	if user.Role == 0 {

		fmt.Println("admin logged in")

		util.Ctx = context.WithValue(util.Ctx, "Authorization", token)
		util.Ctx = context.WithValue(util.Ctx, "Admin", user)
		_ = context.WithValue(ctx, "Authorization", token)
		_ = context.WithValue(ctx, "Admin", token)

		checkToken, ok := util.Ctx.Value("Authorization").(string)

		if ok {
			fmt.Println(checkToken)
		}

	} else {
		fmt.Println(username + " logged in")

		ctx = context.WithValue(ctx, "Authorization", token)
		checkToken, ok := ctx.Value("Authorization").(string)

		if ok {
			fmt.Println(checkToken)
		}
	}

	return token, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, input model.ResetPasswordInput) (string, error) {
	panic(fmt.Errorf("not implemented: ResetPassword - resetPassword"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
