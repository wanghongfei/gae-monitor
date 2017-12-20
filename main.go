package main

import (
	"fmt"
	"net/http"
	"gaemonitor/exposure"
)

func main() {
	http.HandleFunc("/", exposure.ExposureHandler)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if nil != err {
		fmt.Println(err)
	}
}
