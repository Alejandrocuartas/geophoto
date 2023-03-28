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
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

const (
	defaultPort     = "8080"
	websocketPath   = "/ws"
	readBufferSize  = 1024
	writeBufferSize = 1024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  readBufferSize,
	WriteBufferSize: writeBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Failed to upgrade WebSocket connection: %v", err)
			return
		}

		// Handle the WebSocket connection
		go handleWebSocketConnection(conn, hub)
	})
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func handleWebSocketConnection(conn *websocket.Conn, hub *sockets.Hub) {
	defer conn.Close()
	wsClient := sockets.NewClient(hub, conn)
	go wsClient.Serve()
}
