package driver

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	//"errors"
	"github.com/alknopfler/Gologger/gologger"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/datamodel"
	"time"
	"errors"
)

//MongoDB struct
type MongoDB struct {
	session    *mgo.Session
	database   string
	collection string
}

//Initialize mongodb  implementation
func (mdb *MongoDB) Initialize(c *database.ConnectionDB) error {
	if c.DbIpaddress == ""  || c.DbProto == "" || c.DbName == "" || c.DbPort == "" || c.DbCollection == "" {
		gologger.Print("ERROR", 1, "Empty value retrieved", "mongodb.go")
		return errors.New("Empty values retrieved")
	}

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{c.DbIpaddress},
		Timeout:  10 * time.Second,
		Database: c.DbName,
		Username: c.DbUsername,
		Password: c.DbPassword,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return err
	}

	session.SetMode(mgo.Monotonic, true)
	mdb.session = session
	mdb.database = c.DbName
	mdb.collection = c.DbCollection

	return nil
}

//InsertEntity mongodb implementation
func (mdb *MongoDB) InsertEntity(i *datamodel.Information) error {
	c := mdb.session.DB(mdb.database).C(mdb.collection)
	err := c.Insert(i)
	if err != nil {
		gologger.Print("ERROR", 2, "Error inserting in mongo", "mongodb.go")
		return err
	}
	return nil
}

//GetEntity mongodb implementation
func (mdb *MongoDB) GetEntity(field, searchItem string) (result []datamodel.Information, err error) {
	c := mdb.session.DB(mdb.database).C(mdb.collection)

	err = c.Find(bson.M{field: searchItem}).All(&result)
	if err != nil {
		gologger.Print("ERROR", 3, "Error searching in mongo", "mongodb.go")
	}
	return result, err
}

//IsNew mongodb implementation
func (mdb *MongoDB) IsNew(field,searchItem string) bool {
	c := mdb.session.DB(mdb.database).C(mdb.collection)
	i, err := c.Find(bson.M{field: searchItem}).Count()
	if err != nil {
		gologger.Print("ERROR", 4, "Error isnew in mongo", "mongodb.go")
		return false
	}
	if i == 0 {
		return true
	}
	return false
}
