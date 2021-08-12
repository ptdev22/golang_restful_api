package main

import (
	Employee_controllers "golang_restful_api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// get by id
	r.HandleFunc("/employee", Employee_controllers.Index).Methods("GET")
	r.HandleFunc("/employee/{id:[0-9]+}", Employee_controllers.Show).Methods("GET")
	r.HandleFunc("/employee", Employee_controllers.Create).Methods("POST")
	r.HandleFunc("/employee/{id:[0-9]+}", Employee_controllers.Update).Methods("PUT")
	r.HandleFunc("/employee/{id:[0-9]+}", Employee_controllers.Delate).Methods("DELETE")

	http.ListenAndServe("127.0.0.1:8080", r)
}
