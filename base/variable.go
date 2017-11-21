package main
// import "fmt"
// var (
// 		a int
// 		b float32
// )
// func main() {
	
// 	// var x int;
// 	// var f float32 = 1.66;
// 	// var s = "abc"
// 	aa := 123;
// 	println(aa);

// 	// var x, y, z int;
// 	var s, n = "abc", 123;

// 	a, b := 222, "haha";
// 	println(n, s, a, b);

// }

// func test() (int, string){
// 	return 1, "abc";
// }

// func main() {
// 	_, s := test();
// 	println(s);
// }


// func main() {
// 	s := "abc";
// 	println(&s);


// 	s, y := "hehe", 20;
// 	println(s, y)

// 	const a = 123;

// }
const (
	aa = iota
	bb = iota
)

func main() {
	// var a int;
	// var b int32;
	// a = 15;
	// b = a + a;
	// b = b + 5;
	// println(aa, bb);
	// if aa == 0 {
	// 	println("a == 0");
	// }
	// for i := 0; i < 10; i++ {
	// 	println(i);
	// }

	// list := []string{"a", "b", "c"};
	// for k, v := range list {
	// 	println(k, ": ", v);
	// }


	// var arr [3]int;
	// arr[0] = 11;
	// arr[1] = 22;
	// arr[2] = 33;
	// // for k, v := range arr{
	// // 	println(k, "= ", v);
	// // }

	// change(arr);
	// println(arr[0]);

	// // s0 := []int{0, 0};
	// // s1 := append(s0, 2);

	// var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7};
	// var s = make([]int, 6);
	// n1 := copy(s, a[0:]);
	// println(s[5], n1)
	// n2 := copy(s, a[2:]);
	// println(s[5], n2);


	// s := "zizizi";
	// c := []byte(s);
	// c[0] = 'x';
	// s2 := string(c);
	// println(s2);

	// s = "ssss" + s[1:];
	// println(s);

	// dbArr := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{2, 3, 4, 5}};

	// for i := 0; i < len(dbArr); i++ {
	// 	for j := 0; j < len(dbArr[i]); j++ {
	// 		println(dbArr[i][j])
	// 	}
	// }

	// a := [3]int{1,2,3};
	// var b [3]int;
	// b = a;
	// b[1] = 1111;
	// println("b[1]", b[1])
	// println("a[1]", a[1])

	// arr := [10]int{1,2,3,4,5,6,7,8,9,10};

	// var aSlice []int;
	// aSlice = arr[0:4];
	// bSlice := arr[3:6];
	// cSlice := append(bSlice, 22);
	// printSlice(aSlice);
	// printSlice(bSlice);
	// printSlice(cSlice);


	numbers := make(map[string]int);
	numbers["index1"] = 111;
	numbers["index2"] = 222;
	numbers["index3"] = 333;

	for k, v := range numbers {
		println(k, "= ", v);
	}






}

func printSlice(arr []int) {
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}
func change(arr [3]int) {
	arr[0] = 323;
	println(arr[0]);
}
























