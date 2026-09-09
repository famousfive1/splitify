package main

import (
	"expense/api"
	"expense/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Services
	authService := service.NewAuthService()
	groupService := service.NewGroupService()

	// API Server handler
	server := api.NewServer(authService, groupService)

	mw, err := server.CreateMiddleware()
	if err != nil {
		log.Fatalln("error creating middleware:", err)
	}

    r := gin.Default()
	r.Use(mw)

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	api.RegisterHandlers(r, server)

	// And we serve HTTP until the world ends.
	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
