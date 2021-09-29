package main

import (
  "fmt"
  "encoding/json"
)

type dimension struct {
   Height int
   Width  int
} 

type dimension2 struct {
   Height int `json:",omitempty"`
   Width  int `json:",omitempty"`
}

type Dog struct {
 	Breed string
	WeightKg int `json:",omitempty"`
        Size dimension `json:",omitempty"`
        Size2 dimension2 `json:",omitempty"`
        Size3 *dimension `json:",omitempty"`
}

func main() {
        d := Dog{
		Breed: "pug",
	}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}
