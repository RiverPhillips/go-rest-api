package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var port = flag.String("port", "8080", "port to serve on")

func main() {
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

	r.Route("/api/v1/todos", handler.NewTodoHandlerV1().Route())

	log.Print("Starting server on port ", *port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), r); err != nil {
		log.Fatal(err)
	}
}
