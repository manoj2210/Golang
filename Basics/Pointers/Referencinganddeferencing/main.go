package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(*b)

	*b = 20

	fmt.Println(*b, a)
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

}
