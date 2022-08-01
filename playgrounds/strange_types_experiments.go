package main

import (
	"fmt"
)

func main() {
	//  c := Collection{ "a", "b", "c" }
	var c Collection
	c.add("d")
	fmt.Println(c)

	var f Function
	f.Echo().Echo2()
}

type Collection []string

func (c *Collection) add(item string) {
	*c = append(*c, item)
}

type Function func(s string)

func (f *Function) Echo() *Function {
	fmt.Println("hello")
	return f
}
func (f *Function) Echo2() *Function {
	fmt.Println("hello2")
	return f
}

