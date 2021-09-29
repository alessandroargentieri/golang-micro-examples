package main

import "fmt"

type Greeter interface {
   Greet()
}

type GreeterImpl struct { }
func(g GreeterImpl) Greet() {
   fmt.Println("Hello from implemented!")
}

type GreeterMock struct { }
func(g GreeterMock) Greet() {
   fmt.Println("Hello from mock!")
}

////////////////////////////////////////////////


var G Greeter      // pointed interface whose impl will be injected 
var GP = &G        // pointer variable which points to the implementation

func SetGreeter(g Greeter) {  // injecting the implementation
   G = g
}

func main() {
 // SetGreeter(GreeterImpl{})  // 1. inject the implementation
   SetGreeter(GreeterMock{})   // 2. inject the mock implementation

   (*GP).Greet()  // the func main can access the local pointer after injection
}


