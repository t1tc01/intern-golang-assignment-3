package middleware

import (
	"context"
	"net/http"

	"gitlab.com/hedwig-phan/assignment-3/cmd/db"
	"gitlab.com/hedwig-phan/assignment-3/cmd/jwt"
	"gitlab.com/hedwig-phan/assignment-3/ent"
	"gitlab.com/hedwig-phan/assignment-3/ent/user"
	"gitlab.com/hedwig-phan/assignment-3/internal/util"
)

type ContextKey struct {
	Name string
}

var userCtxKey = &ContextKey{"user"}

// func MiddlewareAuthenticateJWT() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		header := c.Get("Authorization")
// 		token := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(header), "Bearer"))

// 		fmt.Println("Token: " + token)
// 		// Allow unauthenticated users in
// 		if header == "" {
// 			return c.Next()
// 		}
// 		//
// 		// Validate JWT token
// 		tokenStr := token
// 		username, err := jwt.ParseToken(tokenStr)
// 		fmt.Println(username)
// 		if username == "" {
// 			// fmt.Println("2222")
// 			return &users.InvalidUserError{}
// 		}

// 		if err != nil {
// 			// fmt.Println("1111")
// 			invalidErr := users.InvalidTokenError{Token: username}
// 			return &invalidErr
// 		}

// 		// Create user and check if username exists in the database
// 		user, err := db.Client.User.Query().Where(user.Username(username)).Only(util.Ctx)
// 		if err != nil {
// 			invalidErr := users.InvalidTokenError{Token: username}
// 			return &invalidErr
// 		}

// 		// Put it in context

// 		_ = context.WithValue(c.Context(), userCtxKey, &user)
// 		c.Context().SetUserValue(userCtxKey, &user)

// 		// Call the next handler with our new context
// 		return c.Next()
// 	}
// }

func MiddlewareAuthenticateJWT() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			// token := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(header), "Bearer"))

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			//Create user and check if username exists in the database
			users, err := db.Client.User.Query().Where(user.Username(username)).Only(util.Ctx)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, users)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *ent.User {
	raw, _ := ctx.Value(userCtxKey).(*ent.User)
	return raw
}

func ForAdminContext(ctx context.Context) *ent.User {
	raw, _ := ctx.Value("Admin").(*ent.User)
	return raw
}
