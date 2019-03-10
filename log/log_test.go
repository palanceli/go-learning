package gollog

import (
	"flag"
	"testing"

	"github.com/golang/glog"
)

/*
参考资料
https://godoc.org/github.com/golang/glog#Infoln
*/

func TestBasicUsage(t *testing.T) {
	defer glog.Flush() // 退出之前Flush残余log数据

	flag.Set("logtostderr", "true") // 打印到stderr
	flag.Parse()
	glog.Info("testing info")

	// 若使用go test -run="BasicUsage"  -v=2会报错：
	// go test: illegal bool flag value 2
	glog.V(2).Infoln("testing V2 info")
	// 如果是普通的应用不存在此问题
}
