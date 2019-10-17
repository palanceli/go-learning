package testjson

import (
	"encoding/json"
	"flag"
	"os"
	"testing"

	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	flag.Set("v", "8")
	flag.Set("logtostderr", "true")
	flag.Set("dump-grpc-meta", "true")
	flag.Set("c", "../config.yml")
	flag.Parse()

	defer glog.Flush()

	os.Exit(m.Run())
}

type PublicKey struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func TestUnmarshal1(t *testing.T) {
	s := `[{"name":"Galaxy Nexus", "price":"3460.00"},{"name":"Galaxy Nexus", "price":"3460.00"}]`

	keys := make([]PublicKey, 0)
	if err := json.Unmarshal([]byte(s), &keys); err != nil {
		glog.Errorf("FAILED to Unmarshal. err=%v", err)
		return
	}
	glog.Infof("%+v", keys)
}

type TablePrivilege struct {
	TableName  string   `json:"TableName"`
	FieldNames []string `json:"FieldNames"`
}

func TestUnarshal2(t *testing.T) {
	s := `[{"TableName":"artist", "FieldNames":["birth", "date-death-birth"]}]`

	data := make([]*TablePrivilege, 0)
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		glog.Errorf("FAILED to Unmarshal. err=%v", err)
		return
	}
	glog.Infof("%v", data)
}

// UserPrivilegeItem 输入参数
type UserPrivilegeItem struct {
	UserID int64             `json:"UserId"`
	Tables []*TablePrivilege `json:"Tables"`
}

func TestUnarshal3(t *testing.T) {
	s := `[{"UserId":24, "Tables":[{"TableName":"artist", "FieldNames":["country", "date-death-birth"]}]}, 
	{"UserId":25, "Tables":[{"TableName":"artist", "FieldNames":["deathyear", "birthyear"]}]}
	]`

	data := make([]*UserPrivilegeItem, 0)
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		glog.Errorf("FAILED to Unmarshal. err=%v", err)
		return
	}
	glog.Infof("%v", data)
}

func TestUnarshal4(t *testing.T) {
	s := `{"UserId":1, "Tables":[{"TableName":"artist", "FieldNames":["birth", "date-death-birth"]}]}`

	data := &UserPrivilegeItem{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		glog.Errorf("FAILED to Unmarshal. err=%v", err)
		return
	}
	glog.Infof("%v", data)
}
