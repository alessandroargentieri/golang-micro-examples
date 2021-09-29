package main

import (
   "example/first"
   second "example/second/sub"
   "fmt"
)

func main() {

   fmt.Println(first.First())
   fmt.Println(second.Second())
}

