package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/GAZIMAGomeDDD/url-shortener/internal/handler"
	"github.com/GAZIMAGomeDDD/url-shortener/internal/server"
	"github.com/GAZIMAGomeDDD/url-shortener/internal/storage"
	"github.com/GAZIMAGomeDDD/url-shortener/internal/storage/inmemory"
	"github.com/GAZIMAGomeDDD/url-shortener/internal/storage/postgres"
	"github.com/GAZIMAGomeDDD/url-shortener/pkg/database/postgresdb"
)

var (
	store string
	s     storage.StorageIface
)

func init() {
	flag.StringVar(&store, "store", "in-memory", "store name")
}

func Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	switch store {
	case "postgres":
		connString := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"

		db, err := postgresdb.NewDB(ctx, connString)
		if err != nil {
			log.Fatal(err)
		}

		s, err = postgres.NewStore(ctx, db)
		if err != nil {
			log.Fatal(err)
		}
	case "in-memory":
		s = inmemory.NewStore()
	default:
		log.Fatal(fmt.Errorf("incorrect store name"))
	}

	h := handler.NewHandler(s)
	srv := server.NewServer(h.Init())

	log.Printf("starting server")
	go srv.Run()

	<-ctx.Done()

	log.Print("shutting server down")
	srv.Stop()
}
