package apperror

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrNewUser           = errors.New("user_not_found")
	ErrServer            = errors.New("problems_with_server")
	ErrUserAlreadyExists = errors.New("user_already_exists")
)

// Нужно доабвить appcode, которые будут обозначать какую-то ошибку

var (
	ErrorUserNotFound   = NewError(ErrNewUser, 404, "Invalid request.", "US-0001")
	ErrorInternalServer = NewError(ErrServer, 500, "Internal problems.", "SV-0001")
	ErrorUserExist      = NewError(ErrUserAlreadyExists, 409, "Email exist in storage.", "US-0002")
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    string `json:"code,omitempty"`
}

func NewError(err error, httpStatusCode int, statusText, appCode string) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: httpStatusCode,
		StatusText:     statusText,
		AppCode:        appCode,
	}
}

func (e *ErrResponse) Error() string {
	return e.StatusText
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
