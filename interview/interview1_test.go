package interview

import "testing"

// ========================
// ========================
// 下面代码能否通过，为什么？
type Param map[string]interface{}

type Show struct {
	Param
}

func TestInterview01(t *testing.T) {
	s := new(Show)
	s.Param["RMB"] = 10000
}

// 无法通过，原因是map在使用之前没有初始化
// 参考https://blog.golang.org/go-maps-in-action

// ========================
