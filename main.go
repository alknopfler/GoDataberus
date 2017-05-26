package main


import (
	"github.com/swatlabs/GoDataberus/driver"
	"github.com/swatlabs/GoDataberus/database"
	"github.com/swatlabs/GoDataberus/api"
	"net/http"
	"github.com/alknopfler/Gologger/gologger"
	"os"
	"github.com/sirupsen/logrus"
	"github.com/swatlabs/GoDataberus/utils"
	"globaldevtools.bbva.com/CASE/servicio-cache-v2/logger"
)

func init (){
	go gologger.Init(os.Stdout,logrus.InfoLevel)
}

func main() {
	var drv database.Store
	if utils.CheckEnvironment() {
		switch utils.DbDriver {
		case "mongo":
			drv = mongodb.MongoDB{}
		default:
			logger.Print("FATAL", 2, "Error driver not found", "main.go")
		}
		db := database.ConnectionDB{utils.DbProto,utils.DbIpaddr,utils.DbPort,utils.DbName}
		http.ListenAndServe(":8080", api.HandlerController(&drv, db))
	}else{

		logger.Print("FATAL", 3, "Some env vars not Found", "main.go")
	}
}
