package repository

import "github.com/Enthreeka/refactoring/internal/user"

type User struct {
	//.. pool db
}

func NewUser() user.Repository {
	return &User{}
}
