package service

import (
	"expense/db"
)

type GroupService interface {
	Create(string) (int, error)
	Get(int) (db.Group, error)
	Delete(int) (bool, error)
	Update(int, string) (bool, db.Group, error)
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

func (h *groupService) Delete(id int) (bool, error) {
	return h.groupRepo.Delete(id)
}

func (h *groupService) Update(id int, name string) (bool, db.Group, error) {
	out := db.Group{
		Id: id,
		Name: name,
	}

	success, err := h.groupRepo.Update(out)
	return success, out, err
}
