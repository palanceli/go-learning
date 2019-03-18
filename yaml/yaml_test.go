package golyaml

// 参考资料可参见https://godoc.org/gopkg.in/yaml.v2

import (
	"flag"
	"io/ioutil"
	"os"
	"sync/atomic"
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

func loadConfig(confPath string, config interface{}, defaultConfig interface{}) interface{} {
	var globalConf atomic.Value
	if defaultConfig != nil {
		globalConf.Store(defaultConfig)
	}
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		glog.Fatalf("FAILED to read config file %s. err=%v", confPath, err)
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		glog.Fatalf("FAILED to parse config file %s. err=%v", confPath, err)
	}
	globalConf.Store(config)
	return globalConf.Load()
}

func TestLoadConf(t *testing.T) {
	var defaultConf = Student{
		Name: "Unknown",
		Age:  18,
	}
	var config = Student{}
	confPath := "./test.yml"
	conf := loadConfig(confPath, &config, &defaultConf).(*Student)
	glog.Infof("conf: %v", conf)
	glog.Infof("Name: %s", conf.Name)
}

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func TestConfig(t *testing.T) {
	flag.Set("c", "conf.yml")
	flag.Parse()

	Initialize(&Config{})
	config := Get().(*Config)
	glog.Infof("Host:%s, User:%s, Password:%s", config.Host, config.User, config.Password)
}
