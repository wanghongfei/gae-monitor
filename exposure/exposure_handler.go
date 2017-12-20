package exposure

import (
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
)

var imageBuf []byte

func init()  {
	// read gif
	gifFile, err := os.Open("m.gif")
	if nil != err {
		panic(err)
	}
	defer gifFile.Close()

	imageBuf, err = ioutil.ReadAll(gifFile)
	if nil != err {
		panic(err)
	}
}

func HttpHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}

	req.ParseForm()
	sid := req.Form["sid"]
	fmt.Println(sid)

	w.Header().Add("Content-Type", "image/gif")
	w.Write(imageBuf)
}