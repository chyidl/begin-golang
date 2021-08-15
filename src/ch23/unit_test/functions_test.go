package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Square(op int) int {
	return op * op
}

func TestSquare(t *testing.T) {
	// 表格测试
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := Square(inputs[i])
		if ret != expected[i] {
			t.Errorf("inputs is %d, the expected is %d, the actual %d", inputs[i], expected[i], ret)
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Fatal")
	fmt.Println("End")
}

func TestSomething(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 13, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 123, "they should not be equal")
}
