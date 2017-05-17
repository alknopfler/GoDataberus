package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alknopfler/tdd_api_mongodb/database"
	"time"
	"github.com/alknopfler/tdd_api_mongodb/data_model"
)

type MongoDB struct {
	session    *mgo.Session
	database   string
	collection string
}

func (mdb *MongoDB) Initialize(c *database.ConnectionDB) error {

 	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{c.Ipaddress},
		Timeout:  10 * time.Second,
		Database: c.Dbname,
		//Username: AuthUserName,
		//Password: AuthPassword,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	mdb.session = session
	mdb.database = c.Dbname
	mdb.collection = "mycollection"
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



