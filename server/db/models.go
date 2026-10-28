package db

type Group struct {
	Id int
	Name string
}

type User struct {
	Id int
	Username string
	DisplayName string
}

type Login struct {
	Username string
	Salt string
	Password string
}

type GroupMembership struct {
	GroupId int
	UserId int
	Balance int
}

type Transaction struct {
	Id int
	GroupId int
	Description string
	Date string
	Category string
	Amount int
	Split string
}
