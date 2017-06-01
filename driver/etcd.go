package driver

import (
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"github.com/swatlabs/GoDataberus/database"
	"time"
	"github.com/swatlabs/GoDataberus/datamodel"
	"github.com/alknopfler/Gologger/gologger"
	"errors"


)

type etcdInputKey struct {
	root    string `json:"root"`
	key     string `json:"key"`
	value   string `json:"value"`
}

type Etcd struct {
	kapi   client.KeysAPI
}

//Initialize mongodb  implementation
func (e *Etcd) Initialize(c *database.ConnectionDB) error {
	if c.DbIpaddress == ""  || c.DbProto == "" || c.DbPort == "" {
		gologger.Print("ERROR", 1, "Empty value retrieved", "etcd.go")
		return errors.New("Empty values retrieved")
	}
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
	return nil
}

func (e *Etcd) InsertEntity(i *datamodel.Information) error {

	input := new(etcdInputKey)
	input.root=((*i)["root"]).(string)
	input.key=((*i)["key"]).(string)
	input.value=((*i)["value"]).(string)
	_, err := e.kapi.Set(context.Background(), input.root+input.key, input.value, nil)
	if err != nil {
		gologger.Print("ERROR",2,"Error inserting item in ETCD","etcd.go")
		return err
	}
	return nil
}

func (e *Etcd) GetEntity(field, searchItem string) (result []datamodel.Information, err error) {
	resp, err := e.kapi.Get(context.Background(), field+searchItem, nil)
	if err != nil {
		gologger.Print("ERROR",2,"Error inserting item in ETCD","etcd.go")
		return nil,err
	}
	return resp.Node.Value, nil
}

func (e *Etcd) IsNew(field string, searchItem string) bool {
	/*resp, err := e.kapi.Get(context.Background(), e.root, nil)
	if err != nil {
		gologger.Print("ERROR", 4, "Error isnew in mongo", "mongodb.go")
		return false
	}
	if resp.Node.Value == "" {
		return true
	}*/
	return false
}