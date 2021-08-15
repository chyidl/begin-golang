package benchmark

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}

	assert.Equal("12345", buf.String())
}

// Benchmark 开头
func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	// 与性能测试无关代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

// Benchmark 开头
func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	// 与性能测试无关代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
