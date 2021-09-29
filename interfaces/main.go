package main

import (
   "fmt"
)

// define two interfaces
type inner interface {
   Greet() string
}
type outer interface {
   GetInner() inner
}
//**********************************


// define two struct which implement those interfaces


type InnerImpl struct {
   Greeting string
}
func(inn InnerImpl) Greet() string {
   return inn.Greeting
}


type OuterImpl struct {
   Inn inner  //I don't need InnerImpl here, just be generic!!!!
}
func(out OuterImpl) GetInner() inner {  // here either... I return inner interface not InnerImpl struct
   return out.Inn
}


// create two instaces of the structs above

var Inner InnerImpl = InnerImpl{ Greeting: "hello!!!!!" }
var Outer OuterImpl = OuterImpl{ Inn: Inner }

var Inner2 InnerImpl

func setInner2(in Inner) {
  Inner2 = &in
}

func main() {
   fmt.Println(Outer.GetInner().Greet())
  //not possible:  Inner2 = InnerImpl{ Greeting: "hi"}
  setInner2(InnerImpl{Greeting: "hi!"})   

   fmt.Println(Inner2.Greet())
}
