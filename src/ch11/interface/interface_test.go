package interface_test

import "testing"

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct{}

// 方法签名一致
func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct{}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

// 多态
func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)

	writeFirstProgrammer(goProg)
	writeFirstProgrammer(javaProg)
}
