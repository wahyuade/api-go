package router

import (
	"github.com/gorilla/mux"
	c "github.com/wahyuade/api-go/controllers"
)

func Initialize () *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", c.Welcome)

	api := r.PathPrefix("/api").Subrouter()
	auth := api.Path("/auth").Subrouter()

	auth.Methods("POST").HandlerFunc(c.PostLogin)

	return r
}