package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/utils"
)

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