package api

import (
	"context"
	"expense/service"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func (h Server) CreateMiddleware() (gin.HandlerFunc, error) {
	spec, err := GetSpec()
	if err != nil {
		return nil, fmt.Errorf("loading spec: %w", err)
	}

	return middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: newAuthenticator(h.authService),
		},
	}), nil
}

func newAuthenticator(authService service.AuthService) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return authenticate(authService, ctx, input)
	}
}

func authenticate(authService service.AuthService, ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	cook, err := input.RequestValidationInput.Request.Cookie("session_id")
	if err != nil {
		return err
	}

	userId, err := authService.Verify(cook.Value)
	if err != nil {
		return err
	}

	middleware.GetGinContext(ctx).Set("auth_user_id", userId)

	return nil
}

