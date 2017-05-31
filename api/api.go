package api

import (
	"github.com/gorilla/mux"
)

//HandlerController function
func HandlerController() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v0/connections/{dbType}", HandlerCheckConnections).Methods("PUT")
	r.HandleFunc("/v0/databerus/{dbType}/resources/{uuid}", HandlerInsert).Methods("PUT")
	r.HandleFunc("/v0/databerus/{dbType}/resources/{uuid}/fields/{field}/items/{item}", HandlerSearch).Methods("GET")
	r.HandleFunc("/v0/databerus/{dbType}/resources/{uuid}/exists/{field}/items/{item}", HandlerExists).Methods("GET")
	return r
}

