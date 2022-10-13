// https://play.golang.com/p/_Rj7WIAcJuP

// THESE EXAMPLE SHOWS HOW ELASTIC ARE SLICE AND MAPS: EVEN IF THEY'RE nil YOU CAN LOOP OVER THEM WITH NO ERROR, YOU CAN ASK (FOR MAPS ONLY) A SPECIFIC KEY
// WITHOUT GETTING AN ERROR!!!

package main

import (
	"fmt"
)

func main() {

	// ~~~~ EXAMPLE 1: nil slice (not initialised) doesn't give error in for loop
	
	var mySlice []string = nil
	
	for _, v := range mySlice {
	   fmt.Println(v)
	}
	
	// ~~~~ EXAMPLE 2: nil slice (got from a function) doesn't return an error in for loop
	
	var mySlice2 = getNilSlice()
	
	for _, v := range mySlice2 {
	   fmt.Println(v)
	}
	
	// ~~~~ EXAMPLE 3: sub-slice operation [start-index:total-length] gives error if the initial slice is less long than total-length
	
	mySlice = append(mySlice, "one", "two", "three")
	fmt.Println(len(mySlice))
	mySlice = mySlice[0:3]
	// mySlice = mySlice[0:4] // would return an error
	// mySlice[10] // would return an error
	fmt.Println(len(mySlice))
	
	// ~~~~ EXAMPLE 4: nil map (not initialised) return "" if we ask a non-existing element from it (so it doesn't return an error!)
	
	var mappa map[string]string = nil
	el := mappa["ciao"]
	fmt.Println("Element: " + el)
	
	fmt.Println("End!")
}

func getNilSlice() []string {
    return nil
}
