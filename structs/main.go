package main

import "fmt"

type Pen struct {
  brand string
  color string
}
func (pen Pen) initialize() Pen {
  pen.brand = "Bic"
  pen.color = "Blue"
  return pen
} 

func main() {
   pen := Pen{}
   pen.color = "Red"
   pen.initialize()       // it doesn't override attributes of pen
   fmt.Println(pen.brand + " -> " + pen.color)
   pen = pen.initialize() // now it does!
   fmt.Println(pen.brand + " -> " + pen.color)

   pen2 := Pen{}
   penPointer := &pen2
   (*penPointer).initialize()
   fmt.Println((*penPointer).brand + " -> " + (*penPointer).color)
}
