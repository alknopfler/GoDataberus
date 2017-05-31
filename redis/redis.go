package redisDB

import (
	"github.com/garyburd/redigo/redis"
)

//NewRedis function
func NewRedis() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err.Error())
	}
	return conn

}
