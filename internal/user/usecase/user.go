package usecase

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/Enthreeka/refactoring/internal/app/dto"
	"github.com/Enthreeka/refactoring/internal/apperror"
	"github.com/Enthreeka/refactoring/internal/entity"
	"github.com/Enthreeka/refactoring/internal/user/usecase/repository"
	"github.com/Enthreeka/refactoring/pkg/logger"
)

type ServiceUser struct {
	repository *repository.User
	log        *logger.Logger
}

func NewUser(repository *repository.User, log *logger.Logger) *ServiceUser {
	return &ServiceUser{
		repository: repository,
		log:        log,
	}
}

const store = `users.json`

func (s *ServiceUser) SearchUsers() *entity.UserStore {
	s.log.Info("Start of users searching")

	userStore, err := getDataFromFile(store)
	if err != nil {
		s.log.Info("Error to get data from file in SearchUsers: %s", err)
		return nil
	}

	s.log.Info("Search users completed successfully")
	return userStore
}

func (s *ServiceUser) CreateUser(request dto.CreateUserRequest) (string, *entity.User, error) {
	s.log.Info("Start of user creation")

	userStore, err := getDataFromFile(store)
	if err != nil {
		s.log.Info("Error to get data from file in CreateUser: %s", err)
		return "", nil, err
	}

	userStore.Increment++

	for _, user := range userStore.List {
		if user.Email == request.Email {
			return "", nil, apperror.ErrorUserExist
		}
	}

	userInfo := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	id := strconv.Itoa(userStore.Increment)
	userStore.List[id] = userInfo

	b, err := json.Marshal(&userStore)
	if err != nil {
		s.log.Error("Error in usecase with unmarshal struct: %s", err)
		return "", nil, err
	}
	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		s.log.Error("Error in usecase with write in file: %s", err)
		return "", nil, err
	}

	s.log.Info("Create user completed successfully")
	return id, &userInfo, nil
}

func (s *ServiceUser) GetUser() *entity.UserStore {
	s.log.Info("Start of users getting")

	userStore, err := getDataFromFile(store)
	if err != nil {
		s.log.Info("Error to get data from file in GetUser: %s", err)
		return nil
	}

	s.log.Info("Get user completed successfully")
	return userStore
}

func (s *ServiceUser) UpdateUser(request dto.UpdateUserRequest, id string) (*entity.UserStore, error) {
	s.log.Info("Start of user updating")

	userStore, err := getDataFromFile(store)
	if err != nil {
		s.log.Info("Error to get data from file in UpdateUser: %s", err)
		return nil, err
	}

	if _, ok := userStore.List[id]; !ok {
		return nil, apperror.ErrorUserNotFound
	}

	user := userStore.List[id]
	user.DisplayName = request.DisplayName
	userStore.List[id] = user

	b, _ := json.Marshal(&userStore)
	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		s.log.Error("Error in usecase with write in file: %s", err)
		return nil, err
	}

	s.log.Info("Update user completed successfully")
	return userStore, nil
}

func (s *ServiceUser) DeleteUser(id string) error {
	s.log.Info("Start of user deleting")

	userStore, err := getDataFromFile(store)
	if err != nil {
		s.log.Info("Error to get data from file in DeleteUser: %s", err)
		return err
	}

	if _, ok := userStore.List[id]; !ok {
		return apperror.ErrorUserNotFound
	}

	delete(userStore.List, id)

	b, _ := json.Marshal(&userStore)
	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		s.log.Error("Error in usecase with write in file: %s", err)
		return err
	}

	s.log.Info("Delete user completed successfully")
	return nil
}

func getDataFromFile(store string) (*entity.UserStore, error) {
	file, _ := ioutil.ReadFile(store)
	userStore := entity.UserStore{}

	err := json.Unmarshal(file, &userStore)
	if err != nil {
		return nil, err
	}

	return &userStore, err
}
