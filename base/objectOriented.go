package main
import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h Human) sayHi() {
	fmt.Printf("Hi, I'm %s, my phone number is %s\n", h.name, h.phone);
}

func (e Employee) sayHi() {
	fmt.Printf("Hi, I'm %s, my phone number is %s, my company is %s\n", e.name, e.phone, e.company);
}
func main() {
	mark := Student{Human{"Mark", 21, "110"}, "THU"};
	sam := Employee{Human{"Sam", 33, "120"}, "DG"};

	mark.sayHi();
	sam.sayHi();
}






