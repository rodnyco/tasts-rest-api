package main

import (
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/rodnyco/tasks-rest-api/internal/task"
	"log"
	"net/http"
)

func main() {
	hs := &http.Server{Addr: "localhost:8888", Handler: buildHandler()}
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}

func buildHandler() http.Handler {
	router := routing.New()
	api := router.Group("/api")
	api.Use(content.TypeNegotiator(content.JSON))

	task.RegisterHandlers(api)

	return router
}
