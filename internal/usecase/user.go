package usecase

import (
	"encoding/json"
	"github.com/Enthreeka/refactoring/internal/app/dto"
	"github.com/Enthreeka/refactoring/internal/entity"
	"github.com/Enthreeka/refactoring/internal/usecase/repository"
	"github.com/Enthreeka/refactoring/pkg/logger"
	"io/fs"
	"io/ioutil"
	"strconv"
	"time"
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

	userStore, err := s.getDataFromFile(store)
	if err != nil {
		s.log.Info("%s", err, ":Error to get data from file in SearchUsers")
		return nil
	}

	s.log.Info("Search users completed successfully")
	return userStore
}

func (s *ServiceUser) CreateUser(request dto.CreateUserRequest) string {
	s.log.Info("Start of user creation")

	userStore, err := s.getDataFromFile(store)
	if err != nil {
		s.log.Info("Error to get data from file in CreateUser: %s", err)
		return ""
	}

	userStore.Increment++

	userInfo := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.DisplayName,
	}

	id := strconv.Itoa(userStore.Increment)
	userStore.List[id] = userInfo

	b, err := json.Marshal(&userStore)
	if err != nil {
		s.log.Error("Error in usecase with unmarshal struct: %s", err)
		return ""
	}
	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		s.log.Error("Error in usecase with write in file: %s", err)
		return ""
	}

	s.log.Info("Create user completed successfully")
	return id
}

func (s *ServiceUser) GetUser() *entity.UserStore {
	s.log.Info("Start of users getting")

	userStore, err := s.getDataFromFile(store)
	if err != nil {
		s.log.Info("Error to get data from file in GetUser: %s", err)
		return nil
	}

	s.log.Info("Get user completed successfully")
	return userStore
}

func (s *ServiceUser) UpdateUser(request dto.UpdateUserRequest, id string) *entity.UserStore {
	s.log.Info("Start of user updating")

	userStore, err := s.getDataFromFile(store)
	if err != nil {
		s.log.Info("Error to get data from file in UpdateUser: %s", err)
		return nil
	}

	user := userStore.List[id]
	user.DisplayName = request.DisplayName
	userStore.List[id] = user

	b, _ := json.Marshal(&userStore)
	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		s.log.Error("Error in usecase with write in file: %s", err)
		return nil
	}

	s.log.Info("Update user completed successfully")
	return userStore
}

func (s *ServiceUser) DeleteUser(id string) *entity.UserStore {
	s.log.Info("Start of user deleting")

	userStore, err := s.getDataFromFile(store)
	if err != nil {
		s.log.Info("Error to get data from file in DeleteUser: %s", err)
		return nil
	}

	delete(userStore.List, id)

	b, _ := json.Marshal(&userStore)
	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		s.log.Error("Error in usecase with write in file: %s", err)
		return nil
	}

	s.log.Info("Delete user completed successfully")
	return userStore
}

func (s *ServiceUser) getDataFromFile(store string) (*entity.UserStore, error) {
	file, _ := ioutil.ReadFile(store)
	userStore := entity.UserStore{}

	err := json.Unmarshal(file, &userStore)
	if err != nil {
		s.log.Error("Error in usecase with unmarshal struct: %s", err)
		return nil, err
	}

	return &userStore, err
}
