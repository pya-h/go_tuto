package main
import "fmt"

func zero(x *int) {
	*x = 0
}

func main() {
	x := new(int)
	zero(x)
	fmt.Println(*x)
}