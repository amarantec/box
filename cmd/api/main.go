package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/amarantec/box/internal/database"
	"github.com/amarantec/box/internal/handler/routes"
	"github.com/amarantec/box/internal/middleware"
	"github.com/amarantec/box/internal/utils"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	utils.LoadEnv()

	dbConfig, err := utils.BuildDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	Conn, err := database.OpenConnection(ctx, dbConfig)
	if err != nil {
		panic(err)
	}

	defer Conn.Close()

	mux := routes.Router(Conn)
	loggedMux := middleware.LoggerMiddleware(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	fmt.Printf("Server listen on: http://localhost%s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
