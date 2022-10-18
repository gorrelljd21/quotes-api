package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorrelljd21/quotes-starter/gqlgen/graph"
	"github.com/gorrelljd21/quotes-starter/gqlgen/graph/generated"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := (http.ListenAndServe(":"+port, router))
	if err != nil {
		panic(err)
	}
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// put it in context
			ctx := context.WithValue(r.Context(), "API-Key", r.Header.Get("X-Api-key"))

			// and call the next with our new context
			r = r.WithContext(ctx)

			// fmt.Println("in middleware land")
			next.ServeHTTP(w, r)
		})
	}
}
