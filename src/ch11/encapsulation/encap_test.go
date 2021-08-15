package encap_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Logf("e2 is %T", e2)
}

// 定义String方法
// 这种定义方式在实例对应方法被调用时，实例的成员会进行值复制
//func (e Employee) String() string {
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

// 此种方法定义没有copy动作
func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperation(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
