package redisDB

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/swatlabs/GoDataberus/utils"
	"encoding/json"
	"github.com/swatlabs/GoDataberus/database"
)

func TestRetrieveConnectionData(t *testing.T) {

	//create fake uuid and fake connection
	fakeuuid := utils.NewResourceID()
	fakeConn:=database.ConnectionDB{"http","localhost","27017", "tests","","","testCollection"}
	encode, _ := json.Marshal(fakeConn)
	(NewRedis()).Do("LPUSH", fakeuuid, encode)

	//test the function
	result:=RetrieveConnectionData(fakeuuid)

	//asserts
	assert.NotEmpty(t,result)
	assert.NotEmpty(t,result.DbName)
	assert.NotEmpty(t,result.DbIpaddress)
	assert.NotEmpty(t,result.DbPort)
	assert.NotEmpty(t,result.DbProto)

}