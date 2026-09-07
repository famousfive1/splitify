package main

import (
	"expense/api"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func createMiddleware() (gin.HandlerFunc, error) {
	spec, err := api.GetSpec()
	if err != nil {
		return nil, fmt.Errorf("loading spec: %w", err)
	}

	return middleware.OapiRequestValidator(spec), nil
}

func main() {
    r := gin.Default()

	mw, err := createMiddleware()
	if err != nil {
		log.Fatalln("error creating middleware:", err)
	}
	r.Use(mw)

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()
	api.RegisterHandlers(r, server)

	// And we serve HTTP until the world ends.
	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
