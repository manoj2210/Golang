package main

//Only packages with name main gets executed else they remain only
// as libraries. We cannot build or execute them but
// go install will place the package inside the pkg directory of the workspace.

import "fmt"

//fmt is a built-in package it's used for formating text

//Like C main gets executed first
func main() {
	fmt.Println("Hello from main")
}

//If there exists an init it gets executed first after init only main
// gets executed. If multiple inits are there they are executed in the
//order of declaration
func init() {
	fmt.Println("Hello from init1")

}

func init() {
	fmt.Println("Hello from init2")
}
