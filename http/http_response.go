package http

import "net/http"

func BadRequest(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusBadRequest)
}

func InternalServerError(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusInternalServerError)
}

func Created(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusCreated)
}
