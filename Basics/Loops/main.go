package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { //i has local scope only
		fmt.Println(i)
	}

	for {
		if i := 0; i%10 == 0 {
			break
		} else {
			fmt.Println(i)
		}
	}

}

//C like Loops but instead of while for is enough
