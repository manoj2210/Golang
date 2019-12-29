package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"` // instead of Age as key wisdon score is present
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func main() {
	p := person{
		"Manoj", "Kumar", 19,
	}
	fmt.Println(p)

	p.Greeting() //Methods to the struct

	bs, _ := json.Marshal(p) //To json
	fmt.Println(string(bs))

	var p1 person

	json.Unmarshal(bs, &p1)

	fmt.Println(p1)

	json.NewEncoder(os.Stdout).Encode(p1) // Better to use encode and decode
	json.NewDecoder(strings.NewReader(`{"First":"Manoj", "Last":"Kumar", "Age":20}`)).Decode(&p1)

}

// While encoding or Marshal we need the Field name with starting Caps
// TO access it outside the package
