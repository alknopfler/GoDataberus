package api

import (
	"github.com/gorilla/mux"
	"github.com/swatlabs/GoDataberus/utils"
	"net/http"

	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/redis"
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

	dbType, _ := mux.Vars(r)["dbType"]
	drv := utils.GetDriver(dbType)

	if db, err := utils.GetDataFromBody(r); err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
	} else {
		if err = drv.Initialize(&db.Connection); err != nil {
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

	var unencoded database.BodyRequest
	uuid, _ := mux.Vars(r)["uuid"]
	dbType, _ := mux.Vars(r)["dbType"]
	drv := utils.GetDriver(dbType)

	dbconnect, _ := redis.Strings((redisDB.NewRedis()).Do("LRANGE", uuid, 0, -1))
	json.Unmarshal([]byte(dbconnect[0]), &unencoded.Connection)

	if err := drv.Initialize(&unencoded.Connection); err != nil {
		responseWithError(w, http.StatusServiceUnavailable, err.Error())
	} else {
		input, err := utils.GetDataFromBody(r)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, err.Error())
		}
		drv.InsertEntity(&input.Message)
	}
}

//HandlerSearch function
func HandlerSearch(w http.ResponseWriter, r *http.Request) {

	var unencoded database.BodyRequest
	field, _ := mux.Vars(r)["field"]
	uuid, _ := mux.Vars(r)["uuid"]
	dbType, _ := mux.Vars(r)["dbType"]
	item, _ := mux.Vars(r)["item"]
	drv := utils.GetDriver(dbType)

	dbconnect, _ := redis.Strings((redisDB.NewRedis()).Do("LRANGE", uuid, 0, -1))
	json.Unmarshal([]byte(dbconnect[0]), &unencoded.Connection)
	if err := drv.Initialize(&unencoded.Connection); err != nil {
		responseWithError(w, http.StatusServiceUnavailable, err.Error())
	} else {
		result, err := drv.GetEntity(field, item)
		if err != nil {
			responseWithError(w, http.StatusNotFound, err.Error())
		}
		responseWithJSON(w, http.StatusOK, result)
	}
}

//HandlerExists need to refactor (duplicated code)
func HandlerExists(w http.ResponseWriter, r *http.Request) {

	var unencoded database.BodyRequest
	field, _ := mux.Vars(r)["field"]
	uuid, _ := mux.Vars(r)["uuid"]
	dbType, _ := mux.Vars(r)["dbType"]
	item, _ := mux.Vars(r)["item"]
	drv := utils.GetDriver(dbType)

	dbconnect, _ := redis.Strings((redisDB.NewRedis()).Do("LRANGE", uuid, 0, -1))
	json.Unmarshal([]byte(dbconnect[0]), &unencoded.Connection)
	if err := drv.Initialize(&unencoded.Connection); err != nil {
		responseWithError(w, http.StatusServiceUnavailable, err.Error())
	} else {
		isNew := drv.IsNew(field, item)
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
