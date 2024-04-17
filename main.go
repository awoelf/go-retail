package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/awoelf/go-retail/db"
	"github.com/awoelf/go-retail/graph"
	"github.com/awoelf/go-retail/services"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func loadEnv() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file.")
	}
}

func (app *Application) Serve() error {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", app.Config.Port)
	return http.ListenAndServe(":"+app.Config.Port, nil)
}

func main() {
	loadEnv()

	dsn := os.Getenv("DSN")
	dbConn, err := db.Connect(dsn)
	if err != nil {
		log.Panic(err)
	}
	services.Register(dbConn.Client)

	defer dbConn.Client.Close()
	
	cfg := Config {
		Port: os.Getenv("PORT"),
	}
	app := &Application {
		Config: cfg,
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}