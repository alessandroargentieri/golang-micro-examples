package main

import (
  "fmt"
)


//****************************************
type container struct {
  contained *string 
}

// variable
var containedValue = "ciao"
//struct with pointer inside
var myContainer container = container{ contained: &containedValue }
//pointer of struct with pointer inside
var pointerContainer *container = &myContainer
//***************************************


func main() {
   fmt.Println(*myContainer.contained)
   // you don't need to use * double times
   fmt.Println(*pointerContainer.contained)
}
