package api

import (
	"context"
	"expense/service"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)


type Server struct{
	authService service.AuthService
	groupService service.GroupService
}

func NewServer(authService service.AuthService, groupService service.GroupService) *Server {
	return &Server{
		authService: authService,
		groupService: groupService,
	}
}

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


var _ ServerInterface = (*Server)(nil)

// (POST /api/auth/login)
func (h *Server) Login(c *gin.Context) {
	var input LoginJSONRequestBody
	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, "Cannot parse JSON")
		return
	}

	token, err := h.authService.Login(input.Username, input.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while logging in")
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"session_id",
		token,
		60*60*24*3,
		"/api",
		"",
		true,
		true,
	)
	c.Status(http.StatusOK)
}


// (GET /api/groups)
func (h *Server) GetMyGroups(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

// (POST /api/groups)
func (h *Server) CreateGroup(c *gin.Context) {
	var input CreateGroupJSONRequestBody
	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, "Cannot parse JSON")
		return
	}

	id, err := h.groupService.Create(input.Name)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while creating group")
		return
	}

	c.JSON(http.StatusOK, Group{
		Id: id,
		Name: input.Name,
	})
}

// (DELETE /api/groups/{id})
func (h *Server) DeleteGroup(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

// (GET /api/groups/{id})
func (h *Server) GetGroup(c *gin.Context, id int) {
	out, err := h.groupService.Get(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Cannot get group")
	}

	c.JSON(http.StatusOK, Group{
		Id: out.Id,
		Name: out.Name,
	})
}

// (PUT /api/groups/{id})
func (h *Server) UpdateGroup(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

