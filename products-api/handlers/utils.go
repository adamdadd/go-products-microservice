package handlers

import (
	"net/http"
)

type CrudController interface {
	GetAll(rw http.ResponseWriter, r *http.Request )
	Add(rw http.ResponseWriter, r *http.Request)
	Update(rw http.ResponseWriter, r *http.Request)
	Delete(rw http.ResponseWriter, r *http.Request)
}
