package main

import (
	"context"
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo/handler"
	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
)

var port = flag.String("port", "8080", "port to serve on")

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pqx", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	flag.Parse()

	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Compress(5),
		middleware.Timeout(10*time.Second),
	)

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	todoRepo := repository.NewTodoRepository(dbpool)

	r.Route("/api/v1/todos", handler.NewTodoHandlerV1(
		todoRepo,
	).Route())

	log.Print("Starting server on port ", *port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), r); err != nil {
		log.Fatal(err)
	}
}
