package main

import (
	"fmt"
	"time"
)

// CHANNELS

func main2() {

	// create a buffered channel
	// with a capacity of 2.
	ch := make(chan string, 2)
	ch <- "geeksforgeeks"
	ch <- "geeksforgeeks world"

	v1 := <-ch
	v2 := <-ch

	fmt.Println(v1)
	fmt.Println(v2)
}

func mainX() {

	mychan := make(chan int, 1)
	mychan <- 1
	value1 := <-mychan

	fmt.Println(value1)

}

func main() {

	// ~~~~~~~~~~~ UNBUFFERED CHANNELS ~~~~~~~~~~~~~ //

	// IMPORTANT: CHAN WITH BUFFER = 0 IS LIKE AN UNBUFFERED CHAN
	// mychan := make(chan int)
	mychan := make(chan int, 0)

	// IMPORTANT: UNBUFFERED CHAN CAN BE FILLED ONLY IN A SEPARATE GOROUTINE ELSE IT WOULD HAPPEN THIS:
	// mychan <- 0    // fatal error: all goroutines are asleep - deadlock!

	// IMPORTANT: I CAN FILL THE UNBUFFERED CHAN WITH SO MANY ITEMS, EVEN IF THERE IS NO READER
	// THE INSERTION WILL REMAIN UPPENDED IN THE SPECIFIC GO-ROUTINE UNTIL SPACE IS MADE BY READING.
	// NO DEADLOCK EVEN IF THERE IS NO BUFFER IN THE CHAN!
	go func() { mychan <- 1 }()
	go func() { mychan <- 2 }()
	go func() { mychan <- 3 }()
	go func() { mychan <- 4 }()
	go func() { mychan <- 5 }()
	go func() { mychan <- 6 }()
	go func() { mychan <- 7 }()
	go func() { mychan <- 8 }()
	go func() { mychan <- 9 }()
	go func() { mychan <- 10; mychan <- 11; mychan <- 12 }()

	value := <-mychan
	fmt.Println(value)

	value = <-mychan
	fmt.Println(value)

	time.Sleep(3 * time.Second)
	value = <-mychan
	fmt.Println(value)

	// ~~~~~~~~~~~ BUFFERED CHANNELS ~~~~~~~~~~~~~ //

	// BUFFERED CHANNELS DON'T LOCK THE THREAD UNTIL THEY REACH THE CAPACITY
	mychan2 := make(chan int, 3)

	mychan2 <- 1
	mychan2 <- 2
	mychan2 <- 3
	// mychan2 <- 4  // fatal error: all goroutines are asleep - deadlock! We get over the capacity limit!
	go func() { mychan2 <- 4 }() // gets over the capacity limit but NO PROBLEM! is into another goroutine!

	fmt.Println(<-mychan2) // 1
	fmt.Println(<-mychan2) // 2
	fmt.Println(<-mychan2) // 3
	fmt.Println(<-mychan2) // 4 async
	// fmt.Println(<-mychan2) fatal error: all goroutines are asleep - deadlock!
	go func() { fmt.Println(<-mychan2) }() // 5
	go func() { fmt.Println(<-mychan2) }() // will listen with no results; if it was synchronous would have caused a deadlock of the main thread
	mychan2 <- 5
	time.Sleep(3 * time.Second)
	fmt.Println("Finished!")

	// ~~~~~~~~~~~ FOR LOOP AND CLOSING A CHANNEL ~~~~~~~~~~~~~~ //

	chan3 := make(chan string, 3)

	chan3 <- "a"
	chan3 <- "b"
	chan3 <- "c"

	fmt.Println(<-chan3)
	fmt.Println(<-chan3)
	fmt.Println(<-chan3)

	chan3 <- "d"
	chan3 <- "e"
	chan3 <- "f"

	close(chan3)

	// IMPORTANT: the for goes in deadlock if you don't close the channel first!!!!
	for v := range chan3 {
		fmt.Println("Read value in for: ", v)
	}

	// IMPORTANT: to check if a channel is closed:
	v2, ok := <-chan3
	if !ok {
		fmt.Println("chan3 is closed!")
	} else {
		fmt.Println(v2)
	}

	// ----
	chan4 := make(chan int)
	go func() { chan4 <- 80; chan4 <- 81; chan4 <- 82; chan4 <- 83 }()
	go func() {
		// IMPORTANT: if set in a GOROUTINE is good because it is a continuous asynchronous reading task, it remains attached to the open channel
		for v := range chan4 {
			fmt.Println("Read value in for (async): ", v)
		}
	}()
	go func() { chan4 <- 84; chan4 <- 85; chan4 <- 86; chan4 <- 87 }()

	time.Sleep(3 * time.Second)
	fmt.Println("Finished!")

}

