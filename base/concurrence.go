package main

import (
	"fmt"
	// "runtime"
)

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		runtime.Gosched();
// 		fmt.Println(s);
// 	}
// }

// func main() {
// 	go say("world");
// 	say("hello");
// }

// func sum(a []int, c chan int) {
// 	total := 0;
// 	for _, v := range a {
// 		total += v;
// 	}
// 	c <- total;
// }

// func main() {
// 	a := []int{2, 6, 9, 1, 22, 5};

// 	c := make(chan int);
// 	go sum(a[:len(a)/2], c);
// 	go sum(a[len(a)/2:], c);
// 	go sum(a[:], c);
// 	x, y ,z := <-c, <-c, <-c;


// 	fmt.Println(x, y, z)
// }



func main() {
	c := make(chan int, 2);
	c <- 1;
	c <- 2;

	fmt.Println(<-c);
	fmt.Println(<-c);
}



























