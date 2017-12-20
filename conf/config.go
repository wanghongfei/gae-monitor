package conf

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type appConfig struct {
	BrokerList		string	`json:"broker_list"`

	ListenHost		string	`json:"listen_host"`
	ListenPort		string `json:"listen_port"`
}

var AppConfig *appConfig

func init() {
	f, err := os.Open("config.json")
	if nil != err {
		panic(err)
	}
	defer f.Close()

	buf, _ := ioutil.ReadAll(f)
	AppConfig = new(appConfig)

	err = json.Unmarshal(buf, AppConfig)
	if nil != err {
		panic(err)
	}
}
