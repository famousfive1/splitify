package api

import (
	"expense/service"
	"net/http"

	"github.com/gin-gonic/gin"
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

var _ ServerInterface = (*Server)(nil)

// (POST /api/auth/login)
func (h *Server) Login(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
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

