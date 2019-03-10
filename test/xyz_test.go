package test

import (
	"flag"
	"os"
	"testing"

	"github.com/golang/glog"
)

/*
参考文档
https://go-zh.org/cmd/go/#hdr-Test_packages
*/

// 可以在该函数前后添加setup 和 teardown的内容
func TestMain(m *testing.M) {
	defer glog.Flush()
	flag.Set("logtostderr", "true")
	flag.Set("v", "0")
	flag.Parse()
	ret := m.Run()
	os.Exit(ret)
}

// 功能测试用例
func TestXYZAdd(t *testing.T) {
	result := XYZAdd(1, 2)
	if result != 3 {
		t.Errorf("Testing XYZAdd(1, 2), Expect:3, actual:%d", result)
	}
}

func TestXYZAddWrong(t *testing.T) {
	result := XYZAddWrong(10, -10)
	if result != 0 {
		t.Errorf("Testing XYZAdd(10, -10), Expect:0, Actual:%d", result)
	}
}
