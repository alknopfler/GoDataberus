package redisDB

import (
	"github.com/garyburd/redigo/redis"
	"github.com/swatlabs/GoDataberus/database"
	"encoding/json"
)

//NewRedis function
func NewRedis() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err.Error())
	}
	return conn

}

func RetrieveConnectionData(uuid string) database.ConnectionDB{
	var connectData database.BodyRequest
	dbconnect, _ := redis.Strings((NewRedis()).Do("LRANGE", uuid, 0, -1))
	json.Unmarshal([]byte(dbconnect[0]), &connectData.Connection)

	return connectData.Connection
}
