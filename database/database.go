package database

import (
	"github.com/swatlabs/GoDataberus/datamodel"
)

//ConnectionDB struct
type ConnectionDB struct {
	DbProto      string `json:"DbProto"`
	DbIpaddress  string `json:"DbIpaddress"`
	DbPort       string `json:"DbPort"`
	DbName       string `json:"DbName"`
	DbUsername   string `json:"DbUsername"`   //this is optional if the type of database requires auth
	DbPassword   string `json:"DbPassword"`   //this is optional if the type of database requires auth
	DbCollection string `json:"DbCollection"` //this is optional just in case that we use mongodb
}

//BodyRequest struct
type BodyRequest struct {
	Connection ConnectionDB          `json:"DBconnection"`
	Message    datamodel.Information `json:"data"`
}

//Store interface
type Store interface {
	Initialize(c *ConnectionDB) error
	InsertEntity(i *datamodel.Information) error
	IsNew(field, searchItem string) bool
	GetEntity(field, searchItem string) ([]datamodel.Information, error)
}
