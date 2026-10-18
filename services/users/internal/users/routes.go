package users

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, c *Controller, auth mux.MiddlewareFunc) {
	protected := r.NewRoute().Subrouter()
	protected.Use(auth)

	protected.HandleFunc("/users/me", c.GetMe).Methods("GET")
}
