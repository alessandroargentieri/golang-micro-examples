package main

import (
  "fmt"
  "time"
  "strconv"
)

var variable string = "VARIABLE"
var pointedValue string = "VALUE"
var pointer *string = &pointedValue

//public func
func SumInts(a int, b int) int {
   return a+b 
} 

//private func
func subInts(a int, b int) int {
   return a-b
}

type Ram struct {
   Model string
   Year  time.Time
}

func (ram Ram) GetModel() string {
   return ram.Model + " " + ram.Year.String()
}

func getPointerValue() string {
   return *pointer
}

func main() {
  fmt.Println(strconv.Itoa(SumInts(13, 34)))
  fmt.Println(strconv.Itoa(subInts(50, 12)))
  
  ram := Ram{
    Model: "DDR SDRAM",
    Year:  time.Date(2019, 6, 5, 11, 35, 04, 0, time.UTC),  //for a new line we need a comma here!
  }
  fmt.Println(ram.GetModel())
}



