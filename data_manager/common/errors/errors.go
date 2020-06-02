package errors

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error      error `json:"-"`
	HttpStatus int   `json:"-"`

	StatusText string `json:"status" example:"Generic error description"`
	ErrorText  string `json:"error,omitempty" example:"Some debug info"`
}

func (res *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, res.HttpStatus)
	return nil
}

func BadRequest(err error) render.Renderer {
	return &ErrorResponse{
		Error:      err,
		HttpStatus: http.StatusBadRequest,
		StatusText: "Bad request",
		ErrorText:  err.Error(),
	}
}

func InternalWithStatus(err error, status string) render.Renderer {
	return &ErrorResponse{
		Error:      err,
		HttpStatus: http.StatusInternalServerError,
		StatusText: status,
		ErrorText:  err.Error(),
	}
}

func Internal(err error) render.Renderer {
	return InternalWithStatus(err, "Internal server error")
}

var NotFound = &ErrorResponse{HttpStatus: http.StatusNotFound, StatusText: "Not found"}
var Unauthorized = &ErrorResponse{HttpStatus: http.StatusUnauthorized, StatusText: "Unauthorized"}
var Forbidden = &ErrorResponse{HttpStatus: http.StatusForbidden, StatusText: "Access denied"}
var EntityTooLarge = &ErrorResponse{HttpStatus: http.StatusRequestEntityTooLarge, StatusText: "File too large"}
