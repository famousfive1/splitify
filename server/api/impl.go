package api

import (
	"splitify/service"
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

