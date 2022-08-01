package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// when you initialize the object you must specify the embedded object
	h := embedder{embedded{3}, "Mario"}

	// when you print it it's clear which are the embedded attributes (you can see it's not at the same level)
	fmt.Println(h)

	// when you set the embedded variable you don't have to explicit that it's embedded
	h.Num = 4

	// when you access the embedded attribute you don't have to explicit that it's embedded
	fmt.Println(h.Num)

	// when you marshal it's not explicit it's embedded
	jsn, _ := json.Marshal(h)
	fmt.Println(string(jsn))

	// also the unmarshalling works pretty easy
	jsnStr := `{"num": 10, "name": "Andrea"}`
	var newH embedder
	_ = json.Unmarshal([]byte(jsnStr), &newH)
	fmt.Println(newH)

}

type embedded struct {
	Num int `json:"num,omitempty"`
}

type embedder struct {
	embedded
	Name string `json:"name,omitempty"`
}

