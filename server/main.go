package main

import (
	"expense/api"
	"expense/service"
	"fmt"
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func createMiddleware(authService service.AuthService) (gin.HandlerFunc, error) {
	spec, err := api.GetSpec()
	if err != nil {
		return nil, fmt.Errorf("loading spec: %w", err)
	}

	return middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: api.NewAuthenticator(authService),
		},
	}), nil
}

func main() {
	// Services
	authService := service.NewAuthService()
	groupService := service.NewGroupService()

	// API Server handler
	server := api.NewServer(authService, groupService)

    r := gin.Default()

	mw, err := createMiddleware(authService)
	if err != nil {
		log.Fatalln("error creating middleware:", err)
	}
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
