package driver

import (
	"errors"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/datamodel"
)

//Fake struct
type Fake struct{}

//Initialize Fake function
func (f *Fake) Initialize(c *database.ConnectionDB) error {
	if c.DbIpaddress == "" || c.DbProto == "" || c.DbName == "" || c.DbPort == "" || c.DbCollection == "" {
		return errors.New("Error Fake Initialize function")
	}
	return nil
}

//InsertEntity Fake function
func (f *Fake) InsertEntity(i *datamodel.Information) error {
	return nil
}

//GetEntity Fake function
func (f *Fake) GetEntity(field, searchItem string) (result []datamodel.Information, err error) {

	result = []datamodel.Information{{"num": "aaa", "strs": "bbb"}}
	return result, nil
}

//IsNew Fake function
func (f *Fake) IsNew(field string, searchItem string) bool {
	return true
}
