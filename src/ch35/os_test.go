package ch35

import (
	"os"
	"testing"
)

func TestOS(t *testing.T) {
	// 创建目录
	t.Log(os.Mkdir("test", 0755))
}
