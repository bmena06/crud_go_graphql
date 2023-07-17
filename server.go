package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bmena06/crud_go/domain/repositories"
	"github.com/bmena06/crud_go/domain/usecase"
	"github.com/bmena06/crud_go/graph"
	"github.com/bmena06/crud_go/infrastructure"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoConnection := infrastructure.Connect()

	userUseCase := usecase.UserUseCase{
		Repository: repositories.UserRepository{
			Client: mongoConnection,
		},
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserUseCase: userUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
