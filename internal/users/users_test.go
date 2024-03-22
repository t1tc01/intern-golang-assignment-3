package users

import (
	"context"
	"log"
	"time"

	"gitlab.com/hedwig-phan/assignment-3/cmd/db"
	"gitlab.com/hedwig-phan/assignment-3/config"
	"gitlab.com/hedwig-phan/assignment-3/ent"
)

func CreateAdmin(c context.Context, username string, email string, password string) (*ent.User, error) {

	current_time := time.Now()

	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	user, err := db.Client.User.Create().
		SetUsername(username).
		SetPassword(hashedPassword).
		SetEmail(email).
		SetRole(config.ROLE_ADMIN).
		SetCreatedAt(current_time).
		SetUpdatedAt(current_time).
		Save(c)

	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}
