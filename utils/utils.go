package utils

import (
	"net/http"
	"github.com/swatlabs/GoDataberus/database"
	"io/ioutil"
	"github.com/alknopfler/Gologger/gologger"
	"encoding/json"

	"github.com/swatlabs/GoDataberus/driver"
	"github.com/satori/go.uuid"
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

func GetDriver(input string) database.Store {
		switch input{
		case "mongo":
			drv := driver.MongoDB{}
			return &drv
		default:
			drv := driver.MongoDB{}
			return &drv
		}
}

func NewResourceID() string {
	return uuid.NewV1().String()

}

