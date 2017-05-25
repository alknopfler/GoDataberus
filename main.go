package main


import (
	"github.com/swatlabs/GoDataberus/driver"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/api"
	"net/http"
)

func main() {
	//the database and driver data, will be passed by environment
	var drv = mongodb.MongoDB{}
	db := database.ConnectionDB{"localhost","service"}

	http.ListenAndServe(":8080", api.HandlerController(&drv,db))
}
