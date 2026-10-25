package db

type Group struct {
	Id int
	Name string
}

type User struct {
	Id int
	Name string
}

type Login struct {
	Username string
	Password string
}
