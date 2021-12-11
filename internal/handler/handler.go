package handler

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func CreateSlug(w http.ResponseWriter, r *http.Request) {

}

func Redirect(w http.ResponseWriter, r *http.Request) {

}
