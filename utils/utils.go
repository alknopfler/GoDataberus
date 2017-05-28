package utils

import (
	"net/http"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/driver"
	"io/ioutil"
	"github.com/alknopfler/Gologger/gologger"
	"encoding/json"
)

func GetDataFromBody(r *http.Request) (database.BodyRequest,error){
	var value database.BodyRequest

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		gologger.Print("ERROR", 1, "Error while reading input JSON", "utils.go")
		return value, err
	}

	err = json.Unmarshal(b, &value)
	if err != nil {
		gologger.Print("ERROR", 2, "Error while unmarshalling input JSON", "utils.go")
		return value, err
	}
	return value,nil
}

func GetDriver(driver string) database.Store{

	switch driver {
	case "mongodb":
		return &driver.MongoDB{}
	default:
		return &driver.MongoDB{}
	}
}
