package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/swatlabs/GoDataberus/utils"

	"encoding/json"
	"github.com/swatlabs/GoDataberus/redis"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

//HandlerController function
func HandlerController() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v0/connections/{dbType}", HandlerCheckConnections).Methods("PUT")
	r.HandleFunc("/v0/databerus/{dbType}/resources/{uuid}", HandlerInsert).Methods("PUT")
	r.HandleFunc("/v0/databerus/{dbType}/resources/{uuid}/fields/{fieldSearch}", HandlerSearch).Methods("GET")
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
			uuid:=utils.NewResourceID()
			encoded,_:=json.Marshal(db.Connection)
			(redisDB.NewRedis()).Do("LPUSH",uuid,encoded)
			responseWithJSON(w,http.StatusCreated,uuid)
		}
	}
}

func HandlerInsert(w http.ResponseWriter, r *http.Request){

	var unencoded database.BodyRequest
	uuid , _ := mux.Vars(r)["uuid"]
	dbType, _ := mux.Vars(r)["dbType"]
	drv:=utils.GetDriver(dbType)

	dbconnect,_:=redis.Strings((redisDB.NewRedis()).Do("LRANGE",uuid,0,-1))
	fmt.Println(dbconnect[0])
	json.Unmarshal([]byte(dbconnect[0]), &unencoded.Connection)
	fmt.Println(unencoded.Connection)

	if err:=drv.Initialize(&unencoded.Connection);  err!=nil{
		responseWithError(w,http.StatusServiceUnavailable,err.Error())
	}else{
		responseWithJSON(w,200,unencoded.Connection)
	}



}

func HandlerSearch(w http.ResponseWriter, r *http.Request){

	/*dbType, _ := mux.Vars(r)["dbType"]
	drv:=utils.GetDriver(dbType)
	search,_ := mux.Vars(r)["fieldSearch"]

	drv.Initialize(&)
*/
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


