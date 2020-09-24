package main

import (
    "os"
    "encoding/json"
    "fmt"
)

type programmer struct { 
  Name string           `json:"nome"` 
  Surname string        `json:"cognome"`  
  Languages []language  `json:"linguaggi"` 
} 

type language struct {
  Name string           `json:"nome"`
  Experience string     `json:"esperienza"`
}

func main() {
  jsonData := []byte(`{
                       "nome":"Daniele",
                       "cognome":"Masera",
                       "linguaggi":[
                          {
                             "nome":"Node.js",
                             "esperienza":"3 anni"
                          },
                          {
                             "nome":"GO",
                             "esperienza":"1 mese"
                          }
                       ]
                    }`)

  if !json.Valid(jsonData) {
    fmt.Printf("JSON is not valid: %s", jsonData)
    os.Exit(1)
  }
  
  var p programmer 
  err := json.Unmarshal(jsonData, &p)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  } 
  fmt.Println(p.Languages[0].Name)

}
