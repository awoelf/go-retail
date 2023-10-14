package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/awoelf/go-retail/server/config"
	"github.com/awoelf/go-retail/server/graph"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	// TO DO: add models
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	port := os.Getenv("PORT")

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, nil)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	cfg := Config {
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConn, err := config.Connect(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	defer dbConn.DB.Close()

	app := &Application {
		Config: cfg,
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
