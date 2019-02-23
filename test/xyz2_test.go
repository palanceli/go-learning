package test

import "testing"

func TestXYZAdd2(t *testing.T) {
	result := XYZAdd(10, 20)
	if result != 30 {
		t.Errorf("Testing XYZAdd(1, 2), Expect:3, actual:%d", result)
	}
}

func TestXYZAddWrong2(t *testing.T) {
	result := XYZAddWrong(100, -100)
	if result != 0 {
		t.Errorf("Testing XYZAdd(10, -10), Expect:0, Actual:%d", result)
	}
}
