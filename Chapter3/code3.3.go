package main
import "fmt"

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("Base Foo")
}
func (base *Base) Bar() {
	fmt.Println("Base Bar")
}

type Child struct {
	Base
}

func (c *Child) Bar() {
	c.Base.Bar();
}

func main() {
	c := Child{};
	c.Bar();
	c.Foo();
}