package main

import (
	"net/http"

	"github.com/DanilKl4/crud-mongo-go/api"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	base := route.PathPrefix("/api").Subrouter()

	base.HandleFunc("/get/{id}", api.GetOne).Methods("GET")
	base.HandleFunc("/get", api.GetAllValues).Methods("GET")
	base.HandleFunc("/create", api.CreateValue).Methods("POST")
	base.HandleFunc("/delete/{id}", api.DeleteById).Methods("DELETE")

	http.ListenAndServe("localhost:5000", route)
}
