package main

import "fmt"

func main() {
	var i int //Initializes i to zero

	var j int = 5 //Declaration and initialization

	// shorthand declaration and initialization,
	// no performance lack because converted in the compile time
	k := 5

	var a, b, c = 1, 2, 3

	fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", j)
	fmt.Printf("%v \n", k)
	fmt.Println(a, b, c)

}
