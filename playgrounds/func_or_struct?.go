package main

import (
	"fmt"
)

func main() {

	// TYPE FUNCTION
	var f Func = func(s string) int { return len(s) }

	// CALLING THE METHOD OF THE TYPE FUNC
	fmt.Println(f.Bellazio("Alessandro"))

	// CALLING THE FUNCTION ITSELF
	fmt.Println(f("Alessandro"))
	
	// ~~~~~~~~~~~~
	
	// STRUCT THAT WANTS TO BE A FUNCTION
	var adder Adder = Adder{ execute: func(a int, b int) int { return a+b; }}
	
	fmt.Println(adder.execute(4, 5))
	
}

// ~~~~~~~~~~~ FUNCTION WHICH WANTS TO BE A STRUCT ~~~~~~~~~~~~~~~~

type Func func(s string) int

func (f *Func) Bellazio(s string) string {
	return fmt.Sprintf("%v ciao", s)
}

// ~~~~~~~~~~~ STRUCT WHICH WANTS TO BE A FUNCTION ~~~~~~~~~~~~~~~~

type Adder struct {
    execute func(a int, b int) int
}

