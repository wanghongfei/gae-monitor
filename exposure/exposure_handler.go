package exposure

import (
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
	"gaemonitor/msg"
	"time"
	"strconv"
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
	sid, exist := req.Form["sid"]
	if exist {
		fmt.Println(sid[0])
		msg.SendMessage("dev-gae-expose", genTimestampString() + "\t" + sid[0])
	}


	w.Header().Add("Content-Type", "image/gif")
	w.Write(imageBuf)
}

func genTimestampString() string {
	ts := time.Now().Unix()
	return strconv.FormatInt(ts, 10)
}