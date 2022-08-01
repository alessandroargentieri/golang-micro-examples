package main

import (
  "net/http"
  "math/rand"
  "log"
  "fmt"
)

/* This playground won't work here because it contains a multiroutine http web server, so it must be copied
   into a main.go file on a computer and run with: go run main.go

   It shows how to return after w.Write() otherwise, let's say an error occurred, it will show the error message and yet the success message.
   Don't forget to return then!
*/

func main() {

   http.HandleFunc("/hello",  func(w http.ResponseWriter, r *http.Request){
     if err := Check(); err != nil {
         w.Write([]byte("An error occurred!")); return
     } 
     
     w.Write([]byte("Hello world!")); return
   })

   log.Fatal(http.ListenAndServe(":8080", nil))
   log.Println("HTTP Server inited on port 8080")
}

func Check() error {
    max := 10
    min := 0
    v := rand.Intn(max-min) + min
    if v <= 5 {
       return fmt.Errorf("Check not passed!")
    }
    return nil
}
