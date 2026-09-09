package service

import "expense/model"

type GroupService interface {
	Create(string) (int, error)
	Get(int) (model.Group, error)
}

type groupService struct {}

func NewGroupService() GroupService {
	return &groupService{}
}

func (h *groupService) Create(name string) (int, error) {
	return 12, nil
}

func (h *groupService) Get(id int) (model.Group, error) {
	return model.Group{ Id: id, Name: "some" }, nil
}

