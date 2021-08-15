package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}

	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("string", i)
		return
	}
	fmt.Println("Unknown Type")
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
