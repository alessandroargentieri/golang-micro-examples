package main

// this playground shows a particular about mutability with pointers.
// 

import (
	"fmt"
)

func main() {
   hello := "Hello"

   bd := BellaDonna{Greet:&hello, Name:"Eloise"}
   fmt.Println("~~~~~ Original values ~~~~~~")
   fmt.Println(*bd.Greet)
   fmt.Println(bd.Name)
 
   bd.ChangeGreet()
   bd.ChangeName()
   fmt.Println("~~~~~ Struct passed as value ~~~~~~")
   fmt.Printf("%s        panic prone approach, pointer didn't change, pointed value yes\n", *bd.Greet)
   fmt.Println(bd.Name)
  
   bd.ChangeGreetWithPointer()
   bd.ChangeNameWithPointer()
   fmt.Println("~~~~~ Struct passed as pointer ~~~~~~")
   fmt.Println(*bd.Greet)
   fmt.Println(bd.Name)

   
}


type BellaDonna struct {
   Greet *string
   Name string
}

// WITHOUT THE POINTER ~~~~~~~~
func (bd BellaDonna) ChangeGreet() {
  newGreet := "Hello2"
  //bd.Greet = &newGreet // NOT WORKING
  *bd.Greet = newGreet // WORKING BUT IF IT'S NULLA IT GOES PANIC
}

// WITH THE POINTER ~~~~~~~~~~~
func (bd *BellaDonna) ChangeGreetWithPointer() {
  newGreet := "Hello3"
  bd.Greet = &newGreet 
  //*bd.Greet = newGreet 
}

// WITHOUT THE POINTER ~~~~~~~~~~~
func (bd BellaDonna) ChangeName() {
  bd.Name = "Claire" 
}

// WITH THE POINTER ~~~~~~~~~~~
func (bd *BellaDonna) ChangeNameWithPointer() {
  bd.Name = "Lorena" 
}

/* printed

~~~~~ Original values ~~~~~~
Hello
Eloise
~~~~~ Struct passed as value ~~~~~~
Hello2        panic prone approach, pointer didn't change, pointed value yes
Eloise
~~~~~ Struct passed as pointer ~~~~~~
Hello3
Lorena

*/
