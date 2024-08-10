package server

import (
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/", welcomeHandler).Methods("GET")
    return router
}
