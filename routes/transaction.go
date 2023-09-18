package routes

import (
	"backend_project/handlers"
	"backend_project/pkg/middleware"
	"backend_project/pkg/mysql"
	"backend_project/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	// r.HandleFunc("/transaction", middleware.Auth(middleware.UploadFile(h.CreateTransaction, "attache"))).Methods("POST")
	r.HandleFunc("/transaction/{id}", h.UpdatesTransaction).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
