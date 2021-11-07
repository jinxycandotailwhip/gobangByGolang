package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	log.Default().Println("gobang backend server starting")
	r.HandleFunc("/api", GobangHandler)
	log.Fatal(http.ListenAndServe(":8090", r))
}
