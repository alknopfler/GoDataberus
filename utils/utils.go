package utils

import (
	"encoding/json"
	"github.com/alknopfler/Gologger/gologger"
	"github.com/swatlabs/GoDataberus/database"
	"io/ioutil"
	"net/http"

	"github.com/satori/go.uuid"
	"github.com/swatlabs/GoDataberus/driver"
)

//GetDataFromBody function
func GetDataFromBody(r *http.Request) (database.BodyRequest, error) {
	var value database.BodyRequest

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		gologger.Print("ERROR", 1, "Error while reading input JSON", "api_utils.go")
		return value, err
	}

	err = json.Unmarshal(b, &value)
	if err != nil {
		gologger.Print("ERROR", 2, "Error while unmarshalling input JSON", "api_utils.go")
		return value, err
	}
	return value, nil
}

//GetDriver function
func GetDriver(input string) database.Store {
	switch input {
	case "mongo":
		drv := driver.MongoDB{}
		return &drv
	case "etcd":
		drv := driver.Etcd{}
		return &drv
	default:
		drv := driver.Fake{}
		return &drv
	}
}

//NewResourceID function
func NewResourceID() string {
	return uuid.NewV1().String()

}
