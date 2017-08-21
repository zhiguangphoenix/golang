package main
import "fmt"

type Integer int;

func (a *Integer) Add(b Integer) {
	*a += b;
}

func main() {
	var a Integer = 1;
	a.Add(2);
	fmt.Println("a = ", a);

	var arr1 = [3]int{1, 2, 3};
	var arr2 = arr1;
	var arr3 = &arr1;

	arr1[1]++;
	fmt.Println(arr2);
	fmt.Println(*arr3);
}