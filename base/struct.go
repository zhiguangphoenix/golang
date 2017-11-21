package main

type person struct {
	name string
	age int
}

func older(p1 person, p2 person) (ageDifference int) {
	if p1.age > p2.age {
		return p1.age - p2.age;
	}else{
		return p2.age - p1.age;
	}
}

type human struct {
	name string
	age int
	weight int
}
type student struct {
	human
	schoolIndex int
}



func main() {
	p1 := person{ name: "jack", age: 12 };
	p2 := person{ name: "tony", age: 33 };

	println(older(p1, p2));

	student1 := student{ human{ "aya", 44, 120 }, 1001 };
	println(student1.name, student1.age, student1.weight, student1.schoolIndex);
}








