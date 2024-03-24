create-db:
	docker run --name postgres -p 5432:5432 -e POSTGRES_DB=earthquake_db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres 

run:
	go run ./cmd

init-schema:
	mkdir db
	cd db; mkdir migration;
	migrate create -ext sql -dir db/migration -seq init_schema
	migrate create -ext sql -dir db/migration -seq user_schema

gen-gql:
	go get github.com/99designs/gqlgen@v0.17.45
	go run github.com/99designs/gqlgen generate
