package database

import (
	"github.com/swatlabs/GoDataberus/data_model"
)


type ConnectionDB struct {
	DbProto         string
	DbIpaddress  	string
	DbPort          string
	DbName		string
	DbUsername	string   //this is optional if the type of database requires auth
	DbPassword      string   //this is optional if the type of database requires auth
	DbCollection    string   //this is optional just in case that we use mongodb
}

type BodyRequest struct{
	Connection	ConnectionDB `json:"DBconnection"`
	Message 	data_model.Information `json:"data"`
}

type Store interface {
	Initialize(c *ConnectionDB) error
	InsertEntity(i *data_model.Information) error
	IsNew(field,searchItem string) bool
	GetEntity(field,searchItem string) ([]data_model.Information, error)
}

