package main

import (
	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()

	glog.V(8).Infof("test glog.V")
}
