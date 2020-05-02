package errors

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error      error `json:"-"`
	HttpStatus int   `json:"-"`

	StatusText string `json:"status"`
	ErrorText  string `json:"error,omitempty"`
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
var EntityTooLarge = &ErrorResponse{HttpStatus: http.StatusRequestEntityTooLarge, StatusText: "Too large file"}
