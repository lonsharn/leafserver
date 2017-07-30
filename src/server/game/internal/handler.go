package internal

import (
	"github.com/golang/glog"
	//"github.com/name5566/leaf/gate"
	"reflect"
	"protocal"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&protocal.Test{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*protocal.Test)
	// 消息的发送者
	//a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	glog.Infof("label=%v", m.GetLabel())

	// 给发送者回应一个 Hello 消息
	/*
	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
	*/
}
