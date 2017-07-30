package gamedata

import (
	"github.com/golang/glog"
	"github.com/name5566/leaf/recordfile"
	"reflect"
)

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		glog.Fatalf("%v", err)
	}
	fn := reflect.TypeOf(st).Name() + ".txt"
	err = rf.Read("gamedata/" + fn)
	if err != nil {
		glog.Fatalf("%v: %v", fn, err)
	}

	return rf
}
