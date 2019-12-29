package main

import "fmt"

var x = 42

func main() {
	fmt.Println(x)
	packagelevel()
}

func packagelevel() {
	fmt.Println(x)
}
