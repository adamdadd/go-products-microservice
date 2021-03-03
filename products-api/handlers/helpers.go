package handlers

import (
	"net/http"
)

type CrudHandler interface {
	Get(writer http.ResponseWriter, r *http.Request )
	Add(writer http.ResponseWriter, r *http.Request)
	Update(writer http.ResponseWriter, r *http.Request)
	Delete(writer http.ResponseWriter, r *http.Request)
}
