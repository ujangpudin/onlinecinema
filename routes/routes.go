package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	CategoryRoutes(r)
	FilmRoutes(r)
	// EpisodeRoutes(r)
	TransactionRoutes(r)
	AuthRoutes(r)
}
