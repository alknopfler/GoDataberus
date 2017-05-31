package driver

import (
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"github.com/swatlabs/GoDataberus/database"
	"time"
	"github.com/alknopfler/Gologger/gologger"
	"github.com/swatlabs/GoDataberus/datamodel"
)

type Etcd struct {
	kapi   client.KeysAPI
	root   string
}

//Initialize mongodb  implementation
func (e *Etcd) Initialize(c *database.ConnectionDB) error {
	cfg := client.Config{
		Endpoints:               []string{c.DbProto+"://"+c.DbIpaddress+":"+c.DbPort},
		Transport:               client.DefaultTransport,
		Username:  		 c.DbUsername,
		Password:  		 c.DbPassword,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	cli, err := client.New(cfg)
	if err != nil {
		gologger.Print("ERROR",1,"Error initializing ETCD","etcd.go")
		return err
	}
	e.kapi = client.NewKeysAPI(cli)
	e.root = c.DbCollection
	return nil
}

func (e *Etcd) InsertEntity(i *datamodel.Information) error {
	_, err := e.kapi.Set(context.Background(), e.root, i, nil)
	if err != nil {
		gologger.Print("ERROR",2,"Error inserting item in ETCD","etcd.go")
		return err
	}
	return nil
}

func (e *Etcd) GetEntity(field, searchItem string) (result []datamodel.Information, err error) {
	resp, err := e.kapi.Get(context.Background(), e.root, nil)
	if err != nil {
		gologger.Print("ERROR",2,"Error inserting item in ETCD","etcd.go")
		return nil,err
	}
	return resp, nil
}

func (e *Etcd) IsNew(field string, searchItem string) bool {
	resp, err := e.kapi.Get(context.Background(), e.root, nil)
	if err != nil {
		gologger.Print("ERROR", 4, "Error isnew in mongo", "mongodb.go")
		return false
	}
	if resp.Node.Value == "" {
		return true
	}
	return false
}