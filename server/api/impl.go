package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var _ ServerInterface = (*Server)(nil)

// (POST /api/auth/login)
func (h *Server) Login(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}


// (GET /api/groups)
func (h *Server) GetMyGroups(c *gin.Context) {
	c.JSON(http.StatusOK, GroupsList{
		{12, "asdf"},
		{42, "universe"},
	})
}

// (POST /api/groups)
func (h *Server) CreateGroup(c *gin.Context) {
	var input CreateGroupJSONRequestBody
	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, "Cannot parse JSON")
		return
	}

	c.JSON(http.StatusOK, Group{
		Id: 12,
		Name: input.Name,
	})
}

// (DELETE /api/groups/{id})
func (h *Server) DeleteGroup(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

// (GET /api/groups/{id})
func (h *Server) GetGroup(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

// (PUT /api/groups/{id})
func (h *Server) UpdateGroup(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

