package api

import (
	"encoding/json"
	"net/http"
	"github.com/swatlabs/GoDataberus/utils"
	"github.com/swatlabs/GoDataberus/redis"
)

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

