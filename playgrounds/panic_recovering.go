package main

import (
	"fmt"
	"strings"
)

// RECOVER A PANIC IN A FUNCTION WITH RETURN!

func main() {
	name, surname := SplitName("Alessandro Argentieri")
	fmt.Println(fmt.Sprintf("My name is %s and my surname is %s", name, surname))
}

func SplitName(completeName string) (name string, surname string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from Panic")
			name = "John"
			surname = "Doe"
		}
	}()

	splitted := strings.Split(completeName, " ")
	name = splitted[0]
	surname = splitted[1]
	panic("let's panic a little!")
	return name, surname
}

