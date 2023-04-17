package user

import (
	"github.com/Enthreeka/refactoring/internal/app/dto"
	"github.com/Enthreeka/refactoring/internal/entity"
)

type Service interface {
	SearchUsers() *entity.UserStore
	CreateUser(request dto.CreateUserRequest) (string, *entity.User, error)
	GetUser() *entity.UserStore
	UpdateUser(request dto.UpdateUserRequest, id string) (*entity.UserStore, error)
	DeleteUser(id string) error
}
