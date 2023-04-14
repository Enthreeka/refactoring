package http

import (
	"github.com/Enthreeka/refactoring/internal/apperror"
	"github.com/Enthreeka/refactoring/internal/usecase"
	"github.com/Enthreeka/refactoring/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
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

func (u *User) GeneralPage(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(time.Now().String()))
}

func (u *User) SearchUsers(w http.ResponseWriter, r *http.Request) {

	userStore := u.service.SearchUsers()

	render.JSON(w, r, userStore.List)
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {

	request := usecase.CreateUserRequest{}

	id := u.service.CreateUser(request)

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, apperror.Err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {

	userStore := u.service.GetUser()

	id := chi.URLParam(r, "id")

	render.JSON(w, r, userStore.List[id])
}

func (u *User) UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	request := usecase.UpdateUserRequest{}

	userStore := u.service.UpdateUser(request, id)

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, apperror.Err)
		return
	}

	if _, ok := userStore.List[id]; !ok {
		_ = render.Render(w, r, apperror.ErrorNotFound)
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (u *User) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	userStore := u.service.DeleteUser(id)

	if _, ok := userStore.List[id]; !ok {
		_ = render.Render(w, r, apperror.ErrorNotFound)
		return
	}

	render.Status(r, http.StatusNoContent)
}
