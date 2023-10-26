// PARTIALLY IMPLEMENTED INTERFACES: COOL!
package main

import (
    "fmt"
)

func main() {
   var greetz Greetz = &MyGreetz{}
   greetz.Hi()
   greetz.Hello()
}

type Greetz interface{
   Hi()
   Hello()
}

// ~~~~~~~~~~~~~~~ 
// DEFAULT IMPLEMENTATION
type DefaultGreetz struct{}
func(dg *DefaultGreetz) Hi() { fmt.Println("Hi from default greetz"); }
func(dg *DefaultGreetz) Hello() { fmt.Println("Hello from default greetz"); }
// ~~~~~~~~~~~~~~~

// IMPLEMENTATION WHICH RELIES ON THE DEFAULT ONE FOR THE MISSING METHODS
type MyGreetz struct {
   *DefaultGreetz
}
func (mg *MyGreetz) Hello() {
    fmt.Println("Hello from my greetz");
}

