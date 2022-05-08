package reflecttest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"_"`
}

func TestReflect1(t *testing.T) {
	user := User{ID: 1, Name: "abc", Age: 20}
	userType := reflect.TypeOf(user)
	userValue := reflect.ValueOf(user)
	for i := 0; i < userValue.NumField(); i++ {
		if userValue.Field(i).CanInterface() {
			fmt.Printf("%s %s = %v - tag: %s \n",
				userType.Field(i).Name,
				userType.Field(i).Type,
				userValue.Field(i).Interface(),
				userType.Field(i).Tag.Get("json"))
		}
	}
}

func assignStructByMap(data map[string]interface{},
	structPtr interface{}) (assignedTagNames []string, err error) {
	assignedTagNames = []string{}
	rType := reflect.TypeOf(structPtr)
	rVal := reflect.ValueOf(structPtr)
	if rType.Kind() != reflect.Ptr {
		return assignedTagNames, errors.Wrapf(errors.New("structPtr is not a prt"), "%v", structPtr)
	}
	// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
	rType = rType.Elem()
	rVal = rVal.Elem()

	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		tag := t.Tag.Get("json")
		if v, ok := data[tag]; ok {
			if f.Type().Kind() == reflect.Int {
				f.SetInt(reflect.ValueOf(v).Int())
			} else {
				f.Set(reflect.ValueOf(v))
			}
			assignedTagNames = append(assignedTagNames, tag)
		}
	}
	return assignedTagNames, nil
}

func TestReflect2(t *testing.T) {
	user := &User{ID: 1, Name: "aaa"}

	var err error
	bytes, err := json.Marshal(user)
	assert.Nil(t, err)

	// 把Extension解析成map
	var userInterface interface{}
	err = json.Unmarshal([]byte(bytes), &userInterface)
	assert.Nil(t, err)

	userMap := userInterface.(map[string]interface{})
	user2 := &User{}
	assignStructByMap(userMap, user2)
}
