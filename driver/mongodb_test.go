package driver

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/swatlabs/GoDataberus/datamodel"
	"github.com/swatlabs/GoDataberus/database"
)

//general data and connection test

var mongo MongoDB
var dbc = database.NewConnectionDB("http","localhost","27017", "tests","","","testCollection","")
var dbcError = database.NewConnectionDB("http","1.1.1.1","27088", "tests","","", "testCollection","")

func TestMongoDB_InitializeSuccessfully(t *testing.T) {
	res := mongo.Initialize(dbc)

	assert.NoError(t, res)
	assert.NotEmpty(t, mongo.collection)
	assert.NotEmpty(t, mongo.database)
	assert.NotEmpty(t, mongo.session)
}

func TestMongoDB_InsertEntity(t *testing.T) {
	var info = datamodel.Information{"num":"aaa","strs":"bbb"}

	//drop collection before testing and get session *mgo Mongo
	mongo.Initialize(dbc)
	mongo.session.DB(dbc.DbName).C(dbc.DbCollection).DropCollection()

	err := mongo.InsertEntity(&info)
	assert.NoError(t, err)
}


func TestMongoDB_GetEntity(t *testing.T) {
	mongo.Initialize(dbc)

	res, err := mongo.GetEntity("num", "aaa")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	fmt.Println(res)
}


/*func TestMongoDB_IsNew(t *testing.T) {
	mongo.Initialize(dbc)
	mongo.session.DB(dbc.DbName).C(dbc.DbCollection).DropCollection()

	//the new value should be new
	if !mongo.IsNew("num2", "aaa2") {
		t.Error("Error, item found and it should not!!!!!!")
	}
	//now, we're going to insert new element to test the other side
	//info := datamodel.Information{"spain", "tohu", "template1"}
	info := datamodel.Information{"num2":"aaa2","strs":"bbb"}
	mongo.InsertEntity(&info)
	if mongo.IsNew("num2", "aaa2") {
		t.Error("Error, item not found and it should be")
	}
}*/
