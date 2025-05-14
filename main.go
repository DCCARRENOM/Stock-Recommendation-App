package main

import (
	"context"
	//"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	//"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"os/signal"
	"server/server"
	"syscall"
)

func main() {

	// Read in connection string
	config, err := pgx.ParseConfig("")
	if err != nil {
		log.Fatalf("error parsing connection configuration: %v", err)
	}
	config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
	conn, err := pgx.ConnectConfig(context.Background(), config)
	server.DB = conn
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer server.DB.Close(context.Background())
	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)

	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8081")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	log.Println("Server started")

	<-serverDoneChan

	srv.Shutdown(ctx)
	log.Println("Server stopped")

}
