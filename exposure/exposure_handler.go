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

// 1像素gif图片
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

	// 取sid参数
	req.ParseForm()
	sid, exist := req.Form["sid"]
	if exist {
		fmt.Println(sid[0])
		msg.SendMessage("dev-gae-expose", genTimestampString() + "\t" + sid[0])
	}


	// 返回图片
	w.Header().Add("Content-Type", "image/gif")
	w.Write(imageBuf)
}

func genTimestampString() string {
	ts := time.Now().Unix()
	return strconv.FormatInt(ts, 10)
}