package routes

import (
	"backend_project/handlers"
	"backend_project/pkg/mysql"
	"backend_project/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/category", h.FindCategory).Methods("GET")
	r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET")
	r.HandleFunc("/category", h.CreateCategory).Methods("POST")
	r.HandleFunc("/category/{id}", h.UpdateCategory).Methods("PATCH")
	r.HandleFunc("/category/{id}", h.DeleteCategory).Methods("DELETE")
}
