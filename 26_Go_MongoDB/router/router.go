package router

import (
	"github.com/Prayag2003/connecting_go_and_mongo/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/deleteamovie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
