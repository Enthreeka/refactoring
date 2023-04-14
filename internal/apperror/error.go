package apperror

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

var (
	UserNotFound = errors.New("user_not_found")
)

// Status code - TODO

var (
	ErrInvalidRequest = NewAppError(nil, 400, "Invalid request.", "")
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    string `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func NewAppError(err error, httpStatusCode int, statusText string, appCode string) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: httpStatusCode,
		StatusText:     statusText,
		AppCode:        appCode,
		ErrorText:      err.Error(),
	}
}

func (e *ErrResponse) Error() string {
	return e.ErrorText
}

func (e *ErrResponse) UnWrap() error {
	return e.Err
}

func (e *ErrResponse) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

//func ErrInvalidRequest(err error) render.Renderer {
//	return &ErrResponse{
//		Err:            err,
//		HTTPStatusCode: 400,
//		StatusText:     "Invalid request.",
//		ErrorText:      err.Error(),
//	}
//}
