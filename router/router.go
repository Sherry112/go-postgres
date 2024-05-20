package router

import (
	"go-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/stocks", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stock/{id}", middleware.GetStocks).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	return r
}
