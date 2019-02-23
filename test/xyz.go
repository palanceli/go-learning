package test

// XYZAddFunc是被测试函数
func XYZAdd(x1, x2 int) int {
	return x1 + x2
}

// XYZAddFunc是被测试的错误函数
func XYZAddWrong(x1, x2 int) int {
	return x1 + x2 + 1
}
