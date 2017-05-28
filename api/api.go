package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/utils"
	"fmt"
	"encoding/json"
)

//HandlerController function
func HandlerController() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v0/connections/{dbType}", HandlerCheckConnections).Methods("PUT")
	r.HandleFunc("/v0/databerus/{dbType}", HandlerInsert)
	//r.HandleFunc("/v0/databerus/{databerus_session}/search/{field_string}", HandlerSearch).Methods("GET")
	//r.HandleFunc("/v0/databerus/{databerus_session}/exists/{field_string}", HandlerExists).Methods("GET")
	return r
}


func HandlerCheckConnections(w http.ResponseWriter, r *http.Request) {

		var drv database.Store

		dbType, _ := mux.Vars(r)["dbType"]

		drv=utils.GetDriver(dbType)




		db,err := utils.GetDataFromBody(r)

		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
		}else{
			err=drv.Initialize(&db.Connection)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				w.Write([]byte("Connection to Database: Failed!"))
			}else{
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Connection to Database: Success!"))
			}
		}
}

func HandlerInsert(w http.ResponseWriter, r *http.Request){
	if r.Method == "PUT" {

		/*var drv database.Store

		dbType, _ := mux.Vars(r)["dbType"]
		switch dbType {
		case "mongo":
			drv = &mongodb.MongoDB{}
		default:
			drv = &mongodb.MongoDB{}
		}*/

		bodyReq, _ := utils.GetDataFromBody(r)
		fmt.Println(bodyReq.Message)
	}else{
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


