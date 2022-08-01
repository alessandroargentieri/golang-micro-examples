package main

import (
	"fmt"
	"sync"
	"time"
)

func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am first runner: start...")
	time.Sleep(6 * time.Second)
	fmt.Print("\n...I am first runner: finish")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner: start...")
	time.Sleep(4 * time.Second)
	fmt.Print("\n...I am second runner: finish")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// We are increasing the counter by 2
	// because we have 2 goroutines
	go runner1(wg)
	go runner2(wg)

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
}

func main() {
	// Launching both the runners
	execute()
}

