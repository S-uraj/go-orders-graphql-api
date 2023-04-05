package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/S-uraj/go-orders-graphql-api/database"
	"github.com/S-uraj/go-orders-graphql-api/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// establish connection
	database.ConnectDB()
	// create db
	database.CreateDB()
	// migrate the db with Post model
	database.MigrateDB()

	config := graph.Config{Resolvers: &graph.Resolver{Database: database.DBInstance}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
