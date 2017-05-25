package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/data_model"

)

//HandlerController function
func HandlerController(drv database.Store, db database.ConnectionDB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v0/countries/{country}/services/{serviceid}", func(w http.ResponseWriter, r *http.Request) {
		HandlerServices(w,r,drv,db)
	}).Methods("GET")
	return r
}


func HandlerServices(w http.ResponseWriter, r *http.Request, drv database.Store, db database.ConnectionDB) {
	service, _ := mux.Vars(r)["serviceid"]
	country, _ := mux.Vars(r)["country"]

	//db := database.ConnectionDB{"localhost","services"}
	//driver := mongodb.MongoDB{}

	drv.Initialize(&db)
	info := data_model.Information{country,service,""}

	drv.InsertEntity(&info)
	w.Write([]byte(service))
}
