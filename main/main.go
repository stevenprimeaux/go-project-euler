package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stevenprimeaux/go-project-euler/output"
)

var problems = map[string]interface{}{
	"1": output.PE1,
	"2": output.PE2,
	"3": output.PE3,
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/problem/{problem}", func(w http.ResponseWriter, r *http.Request) {
		problems[mux.Vars(r)["problem"]].(func(http.ResponseWriter, *http.Request))(w, r)
	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}
