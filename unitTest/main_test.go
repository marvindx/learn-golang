package main

import (
	"fmt"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestDemo(t *testing.T) {
	sum := 5050
	t.Log("sum 1 to 100, should equle ", sum)
	rst := 0
	for i := 1; i < 101; i++ {
		rst += i
	}
	if rst == sum {
		t.Log("check right ", checkMark)
	} else {
		t.Error(rst, "check error ", ballotX)
	}
}

// 通常编写一个名称以 `Test` 开头的函数来创建测试。
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` 会报告测试失败的信息，然后继续运行测试。
		// `t.Fail*` 会报告测试失败的信息，然后立即终止测试。
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// 单元测试可以重复，所以会经常使用 *表驱动* 风格编写单元测试，
// 表中列出了输入数据，预期输出，使用循环，遍历并执行测试逻辑。
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// t.Run 可以运行一个 "subtests" 子测试，一个子测试对应表中一行数据。
		// 运行 `go test -v` 时，他们会分开显示。
		testname := fmt.Sprintf("[a:%d,b:%d]", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// 基准测试 测试代码性能 必须以Benchmark开头
func BenchmarkIntMin(b *testing.B) {
	b.ResetTimer() // 重置计时器

	for i := 0; i < b.N; i++ {
		//IntMin(1, 2)
		for i, rst := 1, 0; i < 101; i++ {
			rst += i
		}
	}
	//b.Run("dddd", func(b *testing.B) {
	//	for i, rst := 1, 0; i < 101; i++ {
	//		rst += i
	//	}
	//})
}
