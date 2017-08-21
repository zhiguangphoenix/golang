package main
import "fmt"

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b;
}
func (a *Integer) Add(b Integer) {
	*a += b;
}

func main() {
	var a Integer = 1;
	var b LessAdder = &a;
  fmt.Println(a.Less(22));
  b.Add(10);
  fmt.Println(a);
}