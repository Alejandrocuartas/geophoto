package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Alejandrocuartas/geophoto/database"
	"github.com/Alejandrocuartas/geophoto/graph"
	"github.com/Alejandrocuartas/geophoto/middlewares"
	"github.com/Alejandrocuartas/geophoto/sockets"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

const (
	defaultPort   = "8080"
	websocketPath = "/ws"
)

func main() {
	r := chi.NewRouter()
	r.Use(middlewares.Auth)
	godotenv.Load()
	database.Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	hub := sockets.NewHub()
	go hub.Run()

	r.HandleFunc(websocketPath, func(w http.ResponseWriter, r *http.Request) {
		sockets.SocketsHTTP(w, r, hub)
	})
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
