package db

import (
	"context"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"gitlab.com/hedwig-phan/assignment-3/config"
	"gitlab.com/hedwig-phan/assignment-3/ent"
	"gitlab.com/hedwig-phan/assignment-3/ent/user"
	"golang.org/x/crypto/bcrypt"
)

var Client *ent.Client

func ConnectDb() {
	dbcon, err := ent.Open(dialect.Postgres, "postgresql://postgres:postgres@localhost:5432/earthquake_db?sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run migration.
	ctx := context.Background()
	err = dbcon.Schema.Create(ctx)

	if err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	//Create admin if do not have
	_, err = dbcon.User.Query().Where(user.RoleEQ(0)).Only(ctx)

	if err != nil {
		current_time := time.Now()
		adminPassword := "password"
		hashedPassword, err := HashPassword(adminPassword)
		if err != nil {
			log.Fatal(err)
		}

		_ = dbcon.User.Create().
			SetUsername("admin").
			SetPassword(hashedPassword).
			SetRole(config.ROLE_ADMIN).
			SetEmail("admin@techvify.com.vn").
			SetCreatedAt(current_time).
			SetUpdatedAt(current_time).
			SaveX(ctx)
	}

	Client = dbcon
	log.Println("Connect")
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
