package main

import "fmt"

func main() {
	fmt.Println(greet("Manoj ", "Kumar "))

}
func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
	//We can return multiple values
}
