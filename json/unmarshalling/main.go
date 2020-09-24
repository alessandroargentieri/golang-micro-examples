package main 

import (
  "encoding/json" 
  "fmt"
  "os" 
)

func main() {
  jsonData := []byte(`{"checkNum":123,"amount":200,"category":["gift","clothing"]}`) 
  if !json.Valid(jsonData) {
    fmt.Printf("JSON is not valid: %s", jsonData)
    os.Exit(1)
  }
  //if  we don't know the type
  var v interface{} 
  err := json.Unmarshal(jsonData, &v)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  } 
  fmt.Println(v) 
}
