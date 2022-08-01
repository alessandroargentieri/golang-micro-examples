package main

import (
	"fmt"
)


/*
   EVERYTHING IN GOLANG IS PASSED BY VALUE: even if we pass a pointer, this value is copied and passed as a function arg.
   If you modify the pointer address overriding it, its scope is only valid within the function and not outside.
*/
func main() {
	var str string = "ciao"
	var pnt *string = &str
	printValue(pnt)
	fmt.Println(*pnt)
}

func printValue(pointer *string) {
   other := "hola"
   pointer = &other
   fmt.Println(*pointer);
}

