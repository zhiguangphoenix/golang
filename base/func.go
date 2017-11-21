package main

func max(a int, b int) int {
	if a > b {
		return a;
	}else{
		return b;
	}
}

func add(arg ... int) int {
	res := 0;
	for i := 0; i < len(arg); i++ {
		res += arg[i]
	}
	return res;
}

func maxAndMin(arr []int)(max int, min int) {
	maxRes := arr[0];
	minRes := arr[0];
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			minRes = arr[i];
		}
		if arr[i] > max {
			maxRes = arr[i];
		}
	}

	return maxRes, minRes
}

func add1(num int) int {
	num = num + 1;
	return num;
}

func add2(num *int) int {
	*num = *num + 2;
	return *num;
}
func main() {
	a, b := 1, 2;
	println(max(a, b));
	println(add(1,2,3,4,5,6,7))

	aSlice := []int{1, 2, 12, 10, 6};
	max, min := maxAndMin(aSlice);
	println("max is ", max, " min is ", min);

	originNum := 1;
	add1(originNum);
	println(originNum);
	add2(&originNum);
	println(originNum);

	println("---------------------");
	
	for i := 0; i < 5; i++ {
		defer println(i);
	}
}




