package test

import (
	"testing"
)

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
