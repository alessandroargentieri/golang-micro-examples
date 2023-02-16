// this example shows how the for loop is suitable to receive in input even nil slices/maps without returning any 
error!
// it shows also what happens when we try to get a value that doesn't exist from a slice and from a map, even if 
they're nil!

package main

import (
	"fmt"
)

func main() {

	defer func() {
	    if r := recover(); r!= nil {
               fmt.Println("recovered from", r)
            }
	}()
	
	// iterating over a nil slice

	var elements []string = nil
	for i, e := range elements {
	  fmt.Printf("index: %d, element: %s \n", i, e)
	}
	
	// iterating over a nil map
	
	var mappa map[string]string = nil
	for k, v := range mappa {
	   fmt.Printf("key: %s, value: %s \n", k, v)
	}
	
	// getting a non existent value from a slice returns error "index out of range"
	//_ = elements[1] 
	
	// getting a non existent value from a map (even a nil one!) returns a zero value, so no error!
        _ = mappa["foo"]
	
	fmt.Println("End")
}

