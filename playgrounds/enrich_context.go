package main

import (
	"context"
	"fmt"
)

func main() {
	// create a parent context
	parentCtx := context.Background()

	// create a child context by enriching the parent context with a value
	ctx := context.WithValue(parentCtx, "key", "value")

	// retrieve the value from the child context
	val := ctx.Value("key")

	// print the value
	fmt.Println(val) // output: value
}

