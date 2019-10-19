package reflecttest

import (
	"fmt"
	"reflect"
	"testing"
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
