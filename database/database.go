package database

import (
	"github.com/swatlabs/GoDataberus/datamodel"
)

//ConnectionDB struct
type ConnectionDB struct {
	DbProto      string `json:"Proto"`
	DbIpaddress  string `json:"Ipaddress"`
	DbPort       string `json:"Port"`
	DbName       string `json:"Name"`
	DbUsername   string `json:"Username"`   //this is optional if the type of database requires auth
	DbPassword   string `json:"Password"`   //this is optional if the type of database requires auth
	DbCollection string `json:"Collection"` //this is optional just in case that we use mongodb
	DbRoot       string `json:"Root"`	//this is optional just in case that we use etcd
}

//NewConnectionDB constructor


func NewConnectionDB(proto,ipaddress,port,dbname,dbusername,dbpassword,dbcollection,dbroot string) *ConnectionDB {
	connectionDB := new(ConnectionDB)
	connectionDB.DbProto = proto
	connectionDB.DbIpaddress = ipaddress
	connectionDB.DbPort = port
	connectionDB.DbName = dbname
	connectionDB.DbUsername = dbusername
	connectionDB.DbPassword = dbpassword
	connectionDB.DbCollection = dbcollection
	connectionDB.DbRoot = dbroot
	return connectionDB
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
