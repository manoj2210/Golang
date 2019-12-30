package main

import "fmt"

func Iterate(a string) string {
	for index := 0; index < 3; index++ {
		a += a
	}
	return a
}

func main() {
	fmt.Println(" ")
}
