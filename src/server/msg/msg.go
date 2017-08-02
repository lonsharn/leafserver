package msg

import (
	//"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/protobuf"
	"protocal"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(10, 2, &protocal.Test{})
}
