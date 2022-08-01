package main

import "fmt"

func main() {
	ciao()
}

// IN GO YOU CAN EVEN DEFINE A NAMED STRUCT INTO A FUNCTION OR AN INTERFACE
// IS NOT POSSIBLE TO DEFINE THE STRUCT METHODS
func ciao() {

	type mollusco interface{}

	type calamaro struct {
		name      string
		tentacoli int
	}

	var c mollusco = calamaro{"Pablo", 8}

	fmt.Println(c)
}

