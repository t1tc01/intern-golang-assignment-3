package users

import (
	"context"
	"log"
	"time"

	"gitlab.com/hedwig-phan/assignment-3/cmd/db"
	"gitlab.com/hedwig-phan/assignment-3/ent"
	"gitlab.com/hedwig-phan/assignment-3/ent/user"
	"golang.org/x/crypto/bcrypt"
)

// Create User
func CreateUser(c context.Context, username string, email string, password string, role int) (*ent.User, error) {

	current_time := time.Now()

	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	user, err := db.Client.User.Create().
		SetUsername(username).
		SetPassword(hashedPassword).
		SetEmail(email).
		SetRole(role).
		SetCreatedAt(current_time).
		SetUpdatedAt(current_time).
		Save(c)

	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Authenticate Login
func AuthenticateLogin(c context.Context, username string, password string) bool {
	userLogin := db.Client.User.Query().Where(user.Username(username)).OnlyX(c)
	hash := userLogin.Password

	return CheckPasswordHash(password, hash)
}
