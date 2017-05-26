package utils

import (
	"os"

)

var DbDriver = os.Getenv("DBDRIVER")
var DbProto = os.Getenv("DBPROTO")
var DbIpaddr = os.Getenv("DBIPADDR")
var DbPort = os.Getenv("DBPORT")
var DbName = os.Getenv("DBNAME")

func CheckEnvironment() bool {
	return ! (DbDriver == "" || DbProto == "" || DbIpaddr == "" || DbPort == "" || DbName == "")
}
