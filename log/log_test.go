package gollog

import (
	// "errors"
	"flag"
	"testing"

	"github.com/golang/glog"
	"github.com/pkg/errors"
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

func errorWrapA() (err error) {
	err = errors.New("create error for some reason")
	return
}

func errorWrapB() (err error) {
	err = errorWrapA()
	if err != nil {
		// err = errors.Wrap(err, "FAILED to errorWrapA")
		return
	}
	return
}

func errorWrapC() (err error) {
	err = errorWrapB()
	if err != nil {
		// err = errors.Wrap(err, "FAILED to errorWrapB")
		return
	}
	return
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func TestErrorWrap(t *testing.T) {
	err := errorWrapC()
	glog.Errorf("%+v", err)
	// err = errors.Cause(err)
	// glog.Errorf("%v", err)
	// if err != nil {
	// 	if err, ok := err.(stackTracer); ok {
	// 		glog.Errorf("err=%v", err)
	// 		for _, f := range err.StackTrace() {
	// 			glog.Errorf("%+v\n", f)
	// 		}
	// 	} else {
	// 		glog.Errorf("err=%v", err)
	// 	}
	// }
}
