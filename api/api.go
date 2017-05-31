package api

import (
	"github.com/gorilla/mux"
	"github.com/swatlabs/GoDataberus/utils"
	"net/http"

	"encoding/json"
	"github.com/swatlabs/GoDataberus/redis"
	"github.com/swatlabs/GoDataberus/database"
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

//HandlerCheckConnections function
func HandlerCheckConnections(w http.ResponseWriter, r *http.Request) {

	vars:=retrieveMuxVars(r)

	if db, err := utils.GetDataFromBody(r); err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
	} else {
		if err = vars.drv.Initialize(&db.Connection); err != nil {
			responseWithError(w, http.StatusServiceUnavailable, err.Error())
		} else {
			uuid := utils.NewResourceID()
			encode, _ := json.Marshal(db.Connection)
			(redisDB.NewRedis()).Do("LPUSH", uuid, encode)
			responseWithJSON(w, http.StatusCreated, uuid)
		}
	}
}

//HandlerInsert function
func HandlerInsert(w http.ResponseWriter, r *http.Request) {

	vars:=retrieveMuxVars(r)
	connectData:=redisDB.RetrieveConnectionData(vars.uuid)

	if err := vars.drv.Initialize(&connectData); err != nil {
		responseWithError(w, http.StatusServiceUnavailable, err.Error())
	} else {
		input, err := utils.GetDataFromBody(r)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, err.Error())
		}
		vars.drv.InsertEntity(&input.Message)
	}
}

//HandlerSearch function
func HandlerSearch(w http.ResponseWriter, r *http.Request) {

	vars:=retrieveMuxVars(r)
	connectData:=redisDB.RetrieveConnectionData(vars.uuid)

	if err := vars.drv.Initialize(&connectData); err != nil {
		responseWithError(w, http.StatusServiceUnavailable, err.Error())
	} else {
		result, err := vars.drv.GetEntity(vars.field, vars.item)
		if err != nil {
			responseWithError(w, http.StatusNotFound, err.Error())
		}
		responseWithJSON(w, http.StatusOK, result)
	}
}

//HandlerExists
func HandlerExists(w http.ResponseWriter, r *http.Request) {

	vars:=retrieveMuxVars(r)
	connectData:=redisDB.RetrieveConnectionData(vars.uuid)

	if err := vars.drv.Initialize(&connectData); err != nil {
		responseWithError(w, http.StatusServiceUnavailable, err.Error())
	} else {
		isNew := vars.drv.IsNew(vars.field, vars.item)
		responseWithJSON(w, http.StatusOK, isNew)
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

type reqVars struct {
	dbType 	string
	uuid 	string
	field 	string
	item 	string
	drv     database.Store
}

func retrieveMuxVars(r *http.Request) reqVars{
	var v reqVars
	v.dbType,_ = mux.Vars(r)["dbType"]
	v.uuid,_ = mux.Vars(r)["uuid"]
	v.field,_ = mux.Vars(r)["field"]
	v.item,_ = mux.Vars(r)["item"]
	v.drv = utils.GetDriver(v.dbType)
	return v
}