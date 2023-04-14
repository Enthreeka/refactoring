package http

import (
	"encoding/json"
	"github.com/Enthreeka/refactoring/internal/apperror"
	"github.com/Enthreeka/refactoring/internal/entity"
	"github.com/Enthreeka/refactoring/internal/usecase"
	"github.com/Enthreeka/refactoring/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io/fs"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	service *usecase.ServiceUser
	log     *logger.Logger
}

func NewHandler(service *usecase.ServiceUser, log *logger.Logger) *User {
	return &User{
		service: service,
		log:     log,
	}
}

const store = `users.json`

func (u *User) searchUsers(w http.ResponseWriter, r *http.Request) {

	userStore := u.service.SearchUsers()

	render.JSON(w, r, userStore.List)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile(store)
	s := entity.UserStore{}
	_ = json.Unmarshal(f, &s)

	request := CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, apperror.Err)
		return
	}

	s.Increment++
	u := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.DisplayName,
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = u

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (u *User) getUser(w http.ResponseWriter, r *http.Request) {

	userStore := u.service.GetUser()

	id := chi.URLParam(r, "id")

	render.JSON(w, r, userStore.List[id])
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile(store)
	s := enUserStore{}
	_ = json.Unmarshal(f, &s)

	request := UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, apperror.Err)
		return
	}

	id := chi.URLParam(r, "id")

	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, apperror.Err)
		return
	}

	u := s.List[id]
	u.DisplayName = request.DisplayName
	s.List[id] = u

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusNoContent)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile(store)
	s := UserStore{}
	_ = json.Unmarshal(f, &s)

	id := chi.URLParam(r, "id")

	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, ErrInvalidRequest(UserNotFound))
		return
	}

	delete(s.List, id)

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusNoContent)
}
