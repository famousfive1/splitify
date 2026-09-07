package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Server struct{}

func NewServer() Server {
	return Server{}
}

var _ ServerInterface = (*Server)(nil)


func (Server) GetMyGroups(c *gin.Context) {
	c.JSON(http.StatusOK, GroupsList{
		{12, "asdf"},
		{42, "universe"},
	})
}

// (POST /api/group)
func (Server) CreateGroup(c *gin.Context) {
	var input CreateGroupJSONRequestBody
	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, "Cannot parse JSON")
		return
	}

	if input.Name == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	c.JSON(http.StatusOK, Group{
		Id: 12,
		Name: input.Name,
	})
}

// (DELETE /api/group/{id})
func (Server) DeleteGroup(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

// (GET /api/group/{id})
func (Server) GetGroup(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

// (PUT /api/group/{id})
func (Server) UpdateGroup(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

