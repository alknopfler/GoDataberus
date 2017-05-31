package main

import (
	"github.com/alknopfler/Gologger/gologger"
	"github.com/sirupsen/logrus"
	"github.com/swatlabs/GoDataberus/api"
	"net/http"
	"os"
)

func init() {
	go gologger.Init(os.Stdout, logrus.InfoLevel)
}

func main() {
	err := http.ListenAndServe(":8080", api.HandlerController())
	if err != nil {
		gologger.Print("ERROR", 1, "Error with the Server", "main.go")
	}

}
