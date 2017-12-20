package main

import (
	"net/http"
	"gaemonitor/exposure"
	"gaemonitor/logg"
	"gaemonitor/conf"
)

func main() {
	http.HandleFunc("/", exposure.HttpHandler)

	host := conf.AppConfig.ListenHost
	port := conf.AppConfig.ListenPort

	logg.Logger.Printf("started at %s:%s\n", host, port)
	err := http.ListenAndServe(host + ":" + port, nil)
	if nil != err {
		logg.Logger.Println(err)
	}
}
