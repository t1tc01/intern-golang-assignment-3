package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"gitlab.com/hedwig-phan/assignment-3/cmd/db"
	"gitlab.com/hedwig-phan/assignment-3/graph"
	"gitlab.com/hedwig-phan/assignment-3/middleware"
	"gitlab.com/hedwig-phan/assignment-3/resolver"
	"gitlab.com/hedwig-phan/assignment-3/router"

	_ "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.ConnectDb()

	// app := fiber.New()
	// app.Use(cors.New())

	// app.Use(middleware.MiddlewareAPILogging())

	// router.SetRoutes(app)

	// server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	// app.Get("/", adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/query")))
	// app.All("/query", middleware.MiddlewareAuthenticateJWT(), adaptor.HTTPHandler(server))

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(app.Listen(":" + port))

	app := chi.NewRouter()
	app.Use(middleware.MiddlewareAPILogging())
	app.Use(middleware.MiddlewareAuthenticateJWT())

	router.SetRoutes(app)

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))
	app.Handle("/", playground.Handler("GraphQL playground", "/query"))
	app.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, app))
}
