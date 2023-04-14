package usecase

import (
	"encoding/json"
	"github.com/Enthreeka/refactoring/internal/apperror"
	"github.com/Enthreeka/refactoring/internal/entity"
	"github.com/Enthreeka/refactoring/internal/usecase/repository"
	"github.com/Enthreeka/refactoring/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io/fs"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type ServiceUser struct {
	log        *logger.Logger
	repository *repository.User
}

func NewUser(log *logger.Logger, repository *repository.User) *ServiceUser {
	return &ServiceUser{
		log:        log,
		repository: repository,
	}
}

const store = `users.json`

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (c *CreateUserRequest) Bind(r *http.Request) error { return nil }
func (c *UpdateUserRequest) Bind(r *http.Request) error { return nil }

func (s *ServiceUser) SearchUsers() *entity.UserStore {
	s.log.Info("Start search users")

	userStore := s.getDataFromFile(store)

	s.log.Info("Search users completed successfully")
	return userStore
}

func (s *ServiceUser) createUser() {
	s.log.Info("Start create user")

	userStore := s.getDataFromFile(store)

	request := CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, apperror.Err)
		return
	}

	userStore.Increment++

	userInfo := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.DisplayName,
	}

	id := strconv.Itoa(userStore.Increment)
	userStore.List[id] = userInfo

	b, _ := json.Marshal(&userStore)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (s *ServiceUser) GetUser() *entity.UserStore {
	userStore := s.getDataFromFile(store)

	return userStore
}

func (s *ServiceUser) updateUser() {
	userStore := s.getDataFromFile(store)

	request := UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, apperror.Err)
		return
	}

	id := chi.URLParam(r, "id")

	if _, ok := userStore.List[id]; !ok {
		_ = render.Render(w, r, apperror.ErrorNotFound)
		return
	}

	user := userStore.List[id]
	user.DisplayName = request.DisplayName
	userStore.List[id] = user

	b, _ := json.Marshal(&userStore)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusNoContent)
}

func (s *ServiceUser) deleteUser() {
	userStore := s.getDataFromFile(store)

	id := chi.URLParam(r, "id")

	if _, ok := userStore.List[id]; !ok {
		_ = render.Render(w, r, apperror.ErrorNotFound)
		return
	}

	delete(userStore.List, id)

	b, _ := json.Marshal(&userStore)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusNoContent)
}

func (s *ServiceUser) getDataFromFile(store string) *entity.UserStore {
	file, _ := ioutil.ReadFile(store)
	userStore := entity.UserStore{}

	err := json.Unmarshal(file, &userStore)
	if err != nil {
		s.log.Error("%s", err, ":Error in usecase with unmarshal struct")
		return nil
	}

	return &userStore
}
