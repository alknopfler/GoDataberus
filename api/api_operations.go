package api

import (
	"github.com/swatlabs/GoDataberus/redis"
	"net/http"
	"github.com/swatlabs/GoDataberus/utils"
)

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

//HandlerExists function
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

