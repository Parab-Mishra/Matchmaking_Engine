package api

import (

    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/profiles", CreateProfileHandler).Methods("POST")
	r.HandleFunc("/profiles", GetAllProfiles).Methods("GET")
    r.HandleFunc("/match/{id}", MatchHandler).Methods("GET")
    return r
}
