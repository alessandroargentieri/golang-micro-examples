package main

import (
	"log"
	"os"
)

func main() {
	// first way to use logger: no instantiation
	log.Println("ciao")

	// second way to use logger: with instantiation
	var logger *log.Logger = log.New(os.Stdout, "", log.LstdFlags)
	logger.Println("ciao2")
}

