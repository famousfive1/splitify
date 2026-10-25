package service

import (
	"expense/db"
)

// TODO: Implement Update and Delete for groups
type GroupService interface {
	Create(string) (int, error)
	Get(int) (db.Group, error)
}

type groupService struct {
	groupRepo db.GroupRepo
}

func NewGroupService(groupRepo db.GroupRepo) GroupService {
	return &groupService{
		groupRepo: groupRepo,
	}
}

func (h *groupService) Create(name string) (int, error) {
	return h.groupRepo.Insert(name)
}

func (h *groupService) Get(id int) (db.Group, error) {
	return h.groupRepo.Select(id)
}

