package driver

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/swatlabs/GoDataberus/database"
	"time"
	"github.com/swatlabs/GoDataberus/data_model"
	"github.com/alknopfler/Gologger/gologger"
	"errors"
)

type MongoDB struct {
	session    *mgo.Session
	database   string
	collection string
}

func (mdb *MongoDB) Initialize(c *database.ConnectionDB) error {
	if c.DbIpaddress == "" || c.DbProto == "" || c.DbName == "" || c.DbPort == "" || c.DbCollection == ""{
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

func (mdb *MongoDB) InsertEntity(i *data_model.Information) error {
	c := mdb.session.DB(mdb.database).C(mdb.collection)
	err := c.Insert(i)
	if err != nil {
		fmt.Println("Error while inserting item in mongo")
		return err
	}
	fmt.Println("Item inserted in Mongo")
	return nil
}

func (mdb *MongoDB) GetEntity(field,searchItem string) (result []data_model.Information,err error) {
	c := mdb.session.DB(mdb.database).C(mdb.collection)

	err = c.Find(bson.M{field: searchItem}).All(&result)
	if err != nil {
		fmt.Println("Error while running the query in mongo")
	}

	return result, err
}

func (mdb *MongoDB) IsNew(field string, searchItem string) bool {
	c := mdb.session.DB(mdb.database).C(mdb.collection)
	i, err := c.Find(bson.M{field: searchItem}).Count()
	if err != nil {
		fmt.Println("Error while running the query in mongo")
		return false
	}
	if i == 0 {
		return true
	}
	return false
}



