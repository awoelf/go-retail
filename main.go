package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/awoelf/go-retail/db"
	"github.com/awoelf/go-retail/graph"
	"github.com/awoelf/go-retail/services"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
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

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (app *Application) Serve() error {
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	return r.Run(":"+app.Config.Port)
}

func main() {
	loadEnv()

	dsn := os.Getenv("DB_URL")
	log.Print(dsn)
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