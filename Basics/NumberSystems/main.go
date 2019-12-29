package main

import "fmt"

func main() {
	fmt.Printf("%d \t %b \t %#x \t %x \t %q \n", 42, 42, 42, 42, 42)
}

//C like printf %b-binary
// %#X-hexa(alphabets caps)
// %#x-hexa(alphabets lowercase)
// %x-Normal Hex
// %q-UTF-8
