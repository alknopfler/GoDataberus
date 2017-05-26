package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/data_model"

	"github.com/swatlabs/GoDataberus/driver"
)

//HandlerController function
func HandlerController() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v0/databerus/{dbType}", HandlerDatabase).Methods("PUT")
	return r
}


func HandlerDatabase(w http.ResponseWriter, r *http.Request) {
	dbType, _ := mux.Vars(r)["dbType"]
	var drv database.Store

	switch dbType {
	case "mongo":
		drv = mongodb.MongoDB{}

	}


	drv.Initialize(&db)
	info := data_model.Information{country,service,""}

	drv.InsertEntity(&info)
	w.Write([]byte(service))
}
