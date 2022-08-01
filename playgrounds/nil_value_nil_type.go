package main

import (
	"fmt"
)

func main() {

	niller1 := Niller(false)
	if niller1 == nil {
		fmt.Println("niller1 is nil")
	}
	niller2 := Niller(true)
	if niller2 == nil {
		fmt.Println("niller2 is nil")
	} else {
		fmt.Println("niller2 itself is not nil ...")
		if niller2.(*int) == nil {
			fmt.Println("... but only after casting niller2 we get nil")
		}
	}
}

func Niller(v bool) interface{} {
	var result *int
	if v {
		return result // type not nil, value nil
	}
	return nil // type nil, value nil
}

