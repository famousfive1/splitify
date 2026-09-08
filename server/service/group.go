package service

import "expense/model"

type GroupService interface {
	Create(string) (int, error)
	Get(int) (model.Group, error)
}

type service struct {}

func NewGroupService() GroupService {
	return service{}
}

func (h service) Create(name string) (int, error) {
	return 12, nil
}

func (h service) Get(id int) (model.Group, error) {
	return model.Group{ Id: id, Name: "some" }, nil
}

