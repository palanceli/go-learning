package golyaml

// 参考资料可参见https://godoc.org/gopkg.in/yaml.v2

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

func TestMain(m *testing.M) {
	defer glog.Flush()
	flag.Set("logtostderr", "true")
	flag.Parse()
	ret := m.Run()
	os.Exit(ret)
}

type Score struct {
	ClassName  string `yaml:"class_name"`
	ClassScore uint16 `yaml:"class_score"`
}

type Student struct {
	Name  string  `yaml:"name"`
	Age   uint32  `yaml:"age"`
	Sex   string  `yaml:"sex"`
	Score []Score `yaml:"score"`
}

// 从yml中读入结构体Student
func TestMarshal(t *testing.T) {
	ymlPath := "./test.yml"
	glog.Infof("Reading config file %s ...", ymlPath)
	data, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		glog.Fatalf("FAILED to read config file %s. err = %v", ymlPath, err)
	}

	glog.Infof("Parsing config file %s ...", ymlPath)
	student := Student{}
	err = yaml.Unmarshal(data, &student)
	if err != nil {
		glog.Fatalf("FAILED parse config file %s. err = %v", ymlPath, err)
	}

	glog.Infof("config = %v ", student)
}
