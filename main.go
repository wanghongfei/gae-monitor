package main

import (
	"net/http"
	"gaemonitor/exposure"
	"gaemonitor/logg"
)

func main() {
	http.HandleFunc("/", exposure.HttpHandler)

	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if nil != err {
		logg.Logger.Println(err)
	}
}
