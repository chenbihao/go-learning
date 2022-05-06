package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

// duck type  签名一致
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"asd\")"
}
func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
