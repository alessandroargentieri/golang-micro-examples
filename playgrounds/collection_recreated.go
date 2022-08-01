package main

import (
	"fmt"
)

func main() {
	//  c := Collection{ "a", "b", "c" }
	var c Collection
	c.add("d")
	fmt.Println(c)
}

type Collection []string

func (c *Collection) add(item string) {
	*c = append(*c, item)
}

