package main
import "fmt"

type Rect struct {
	x, y float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height;
}

func main() {
	rect1 := new(Rect);
	rect2 := &Rect{};
	rect3 := &Rect{0, 0, 100, 200};

	fmt.Println(rect1.Area());
	fmt.Println(rect2.Area());
	fmt.Println(rect3.Area());
}