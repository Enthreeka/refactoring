package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Enthreeka/refactoring/internal/app/dto"
	"github.com/Enthreeka/refactoring/internal/apperror"
	"github.com/Enthreeka/refactoring/internal/user/usecase"
	"github.com/Enthreeka/refactoring/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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

func (u *User) SearchUsersHandler(w http.ResponseWriter, r *http.Request) {

	userStore := u.service.SearchUsers()

	if len(userStore.List) == 0 {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]interface{}{
			"error": apperror.ErrorUserNotFound,
		})
		return
	}

	render.JSON(w, r, userStore.List)
}

func (u *User) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	request := dto.CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, user, err := u.service.CreateUser(request)

	if err != nil {
		if err == apperror.ErrorUserExist {
			_ = render.Render(w, r, apperror.ErrorUserExist)
			return
		}
		_ = render.Render(w, r, apperror.ErrorInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id":    id,
		"user_name":  user.DisplayName,
		"user_email": user.Email,
	})
}

func (u *User) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	userStore := u.service.GetUser()

	id := chi.URLParam(r, "id")

	_, ok := userStore.List[id]
	if !ok {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, apperror.ErrorUserNotFound)
		return
	}

	render.JSON(w, r, userStore.List[id])

}

func (u *User) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	request := dto.UpdateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userStore, err := u.service.UpdateUser(request, id)

	if err != nil {
		if err == apperror.ErrorUserNotFound {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, apperror.ErrorUserNotFound)
		}
		return
	}

	if _, ok := userStore.List[id]; !ok {
		_ = render.Render(w, r, apperror.ErrorUserNotFound)
		return
	}

	// if err := render.Bind(r, &request); err != nil {
	// 	_ = render.Render(w, r, apperror.Err)
	// 	return
	// }

	render.JSON(w, r, userStore.List[id])
	render.Status(r, http.StatusNoContent)
}

func (u *User) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := u.service.DeleteUser(id)
	if err != nil {
		if err == apperror.ErrorUserNotFound {
			_ = render.Render(w, r, apperror.ErrorUserNotFound)
			return
		}
		_ = render.Render(w, r, apperror.ErrorInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
