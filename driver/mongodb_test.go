package mongodb

import (
	"testing"
	"github.com/alknopfler/tdd_api_mongodb/database"
	"github.com/stretchr/testify/assert"
	"github.com/alknopfler/tdd_api_mongodb/data_model"

	"fmt"
)

//general data and connection test
var mongo MongoDB
var dbc = database.ConnectionDB{"localhost","test"}
var dbc_error = database.ConnectionDB{"1.1.1.1","test"}


func TestMongoDB_InitializeSuccessfully(t *testing.T) {
	res := mongo.Initialize(&dbc)

 	assert.NoError(t,res)
	assert.NotEmpty(t,mongo.collection)
	assert.NotEmpty(t,mongo.database)
	assert.NotEmpty(t,mongo.session)
}

func TestMongoDB_InitializeError(t *testing.T) {
	res := mongo.Initialize(&dbc_error)

	assert.Error(t,res)
	assert.NotEmpty(t,mongo.collection)
	assert.NotEmpty(t,mongo.database)
	assert.NotEmpty(t,mongo.session)
}

func TestMongoDB_InsertEntity(t *testing.T) {
	info := data_model.Information{"spain","tohu","template1"}

	//drop collection before testing and get session *mgo Mongo
	mongo.Initialize(&dbc)
	mongo.session.DB("test").C("mycollection").DropCollection()

	err:= mongo.InsertEntity(&info)
	assert.NoError(t,err)
}

func TestMongoDB_GetEntity(t *testing.T) {
	mongo.Initialize(&dbc)

	res,err:= mongo.GetEntity("country","spain")
	assert.NoError(t,err)
	assert.NotEmpty(t,res)
	fmt.Println(res)
}

func TestMongoDB_IsNew(t *testing.T) {
	mongo.Initialize(&dbc)
	mongo.session.DB("test").C("mycollection").DropCollection()

	//the new value should be new
	if ! mongo.IsNew("country","spain2"){
		t.Error("Error, item found and it should not!!!!!!")
	}
	//now, we're going to insert new element to test the other side
	info := data_model.Information{"spain","tohu","template1"}
	mongo.InsertEntity(&info)
	if mongo.IsNew("country","spain"){
		t.Error("Error, item not found and it should be")
	}
}


