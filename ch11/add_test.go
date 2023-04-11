package ch11

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	re := add(1, 3)
	if re != 4 {
		t.Errorf("expect %d, actual %d", 3, re)
	}
}

// 执行命令 go test -short
func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip("Short 模式下跳过")
	}
	fmt.Println("yes")
	re := add(1, 3)
	if re != 4 {
		t.Errorf("expect %d, actual %d", 3, re)
	}
}

// 基于 表格驱动测试
func TestAdd3(t *testing.T) {

	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2},
		{12, 12, 24},
		{-9, 8, -1},
		{0, 0, 0},
	}

	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("expect %d actual %d", value.out, re)
		}
	}
}

// Benchmark 性能测试
func BenchmarkAdd(bb *testing.B) {

	var a, b, c int
	a = 100
	b = 20
	c = 120

	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			bb.Errorf("expect %d, actual %d", c, actual)
		}
	}
}
