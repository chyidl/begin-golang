package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// 组合需要重新定义方法
type Dog struct {
	// 内嵌的结构类型
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

/*
func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ", host)
}
*/

func TestDog(t *testing.T) {
	// 不支持LSP
	dog := new(Dog)
	dog.SpeakTo("Chao")
}
