package main

import (
	"fmt"
)

func main() {
        // ONE LINE POINTER WITH 'new'
        p := new(string); *p="ciao"
	fmt.Println(fmt.Sprintf("pointer: %v, value: %s", p, *p))
	
	// ONE LINE POINTER WITH SLICE
	p2 := &[]string{"hello"}[0]
	fmt.Println(fmt.Sprintf("pointer: %v, value: %s", p2, *p2))
}

