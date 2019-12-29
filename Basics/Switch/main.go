package main

import "fmt"

func main() {
	switch "hi" {
	case "hi", "bye": //If any one of the arguements matches with the value
		fmt.Println("hi")
		fallthrough //Implicit break is given to have a fallthrough we use fallthrough
	case "wow", "cow":
		fmt.Println("cow")
	default:
		fmt.Println("default")
	}

}
