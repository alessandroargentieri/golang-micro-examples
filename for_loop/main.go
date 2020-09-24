package main

import ( 
  "fmt"
  "strconv" 
) 

func main() {

   letters := []string{"a", "b", "c"}
   for index, letter := range letters {
      fmt.Println("letter number "+ strconv.Itoa(index) + ": " + letter)
   }

}

