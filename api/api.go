package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/swatlabs/GoDataberus/utils"
	"fmt"
	"encoding/json"
//	"github.com/swatlabs/GoDataberus/database"
//	"github.com/alknopfler/tdd_api_mongodb/driver"
)

//HandlerController function
func HandlerController() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v0/connections/{dbType}", HandlerCheckConnections).Methods("PUT")
	r.HandleFunc("/v0/databerus/{dbType}", HandlerInsert).Methods("PUT")
	//r.HandleFunc("/v0/databerus/{databerus_session}/search/{field_string}", HandlerSearch).Methods("GET")
	//r.HandleFunc("/v0/databerus/{databerus_session}/exists/{field_string}", HandlerExists).Methods("GET")
	return r
}


func HandlerCheckConnections(w http.ResponseWriter, r *http.Request) {

		dbType, _ := mux.Vars(r)["dbType"]
		drv:=utils.GetDriver(dbType)

		if db,err := utils.GetDataFromBody(r); err!= nil{
			responseWithError(w, http.StatusBadRequest, err.Error())
		}else{
			if err=drv.Initialize(&db.Connection);  err!=nil{
				responseWithError(w,http.StatusServiceUnavailable,err.Error())
			}else{
				responseWithJSON(w,http.StatusOK,"Connection to Database: Success!")
			}
		}
}

func HandlerInsert(w http.ResponseWriter, r *http.Request){

		dbType, _ := mux.Vars(r)["dbType"]
		drv:=utils.GetDriver(dbType)
		if db,err := utils.GetDataFromBody(r); err!= nil{
			responseWithError(w, http.StatusBadRequest, err.Error())
		}else{
			if err=drv.Initialize(&db.Connection);  err!=nil{
				responseWithError(w,http.StatusServiceUnavailable,err.Error())
			}else{
				bodyReq, _ := utils.GetDataFromBody(r)
				fmt.Println(bodyReq.Message)
			}
		}


}

func responseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]string{"error": message})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


