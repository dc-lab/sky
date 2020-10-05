package app

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Unauthorized(w http.ResponseWriter) {
	http.Error(w, "No user id specified", http.StatusUnauthorized)
}

type UserNotFound struct{}

func (e *UserNotFound) Error() string {
	return "user not found"
}

type GroupNotFound struct{}

func (e *GroupNotFound) Error() string {
	return "group not found"
}

type PermissionDenied struct{}

func (e *PermissionDenied) Error() string {
	return "permission denied"
}

type WrongPassword struct{}

func (e *WrongPassword) Error() string {
	return "invalid credentials. please try again"
}

type EmptyField struct{}

func (e *EmptyField) Error() string {
	return "required field is empty"
}

func HandleBaseError(w http.ResponseWriter, err error, status int) {
	log.Println(err)
	http.Error(w, err.Error(), status)
}