package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func router() http.Handler {
	router := mux.NewRouter()
	return router
}
