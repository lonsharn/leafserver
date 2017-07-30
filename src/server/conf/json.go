package conf

import (
	"encoding/json"
	"github.com/golang/glog"
	"io/ioutil"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		glog.Fatalf("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		glog.Fatalf("%v", err)
	}
}
