package main

import "fmt"

func main() {

	greeting := int{
		1,2,3,4,5,6,7,8,9
	}

	fmt.Println(greeting[1:2])
	fmt.Println(greeting[:2])
	fmt.Println(greeting[5:])
	fmt.Println(greeting[:])
}

//Slices are like vectors in C++