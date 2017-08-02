package main

import (
	"net"
	"time"
	"fmt"
	"flag"
	"protocal"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"github.com/golang/glog"
)

func init(){
	flag.Parse()
}

func main() {
	defer func(){
		glog.Flush()
	}()

	glog.Info("hello")
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	for {
		msg := &protocal.Test{
			Label:proto.String("hello"),
		}

		data, _:=proto.Marshal(msg)
		// len + data
		m := make([]byte, 2+2+len(data))

		// 默认使用大端序
		binary.BigEndian.PutUint16(m, uint16(len(data)+2))

		// 默认使用大端序
		binary.BigEndian.PutUint16(m[2:], uint16(0))

		copy(m[4:], data)


		// 发送消息
		conn.Write(m)
		fmt.Println("msg len=", len(m))
		time.Sleep(1e9)
	}
	conn.Close()
	time.Sleep(1)
}
