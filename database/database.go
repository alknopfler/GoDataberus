package database

import (
	"golang-course/src/github.com/"
	"github.com/alknopfler/tdd_api_mongodb/data_model"
)


type ConnectionDB struct {
	Ipaddress  	string
	Dbname		string
}

type Store interface {
	Initialize(c *ConnectionDB) error
	InsertEntity(i *data_model.Information) error
	IsNew(field,searchItem string) bool
	GetEntity(field,searchItem string) ([]data_model.Information, error)
}

