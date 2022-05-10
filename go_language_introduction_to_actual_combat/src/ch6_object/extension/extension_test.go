package extension

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("旺旺")
}
func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println("to->", host)
}

func TestDog(t *testing.T) {
	// var dog Pet = new(Dog)
	// 以上无法支持 LSP 李氏替换原则

	dog := new(Dog)
	dog.SpeakTo("chen")
}
