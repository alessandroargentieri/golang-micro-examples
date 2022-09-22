// PARALLEL HTTPCALLS SIMULATION USING GOROUTINES AND A CHANNEL TO GATHER THE RESULTS
// IT'S NOT NECESSARY TO USE A WAITGROUP BECAUSE WE BLOCK THE MAIN GOROUTINE WAITING FOR RESULTS IN THE CHANNEL 
// IF NO BLOCKING READING OPERATION (FROM THE CHANNEL) HAS BEEN INSERTED, THE MAIN GOROUTINE EXECUTION WOULD HAVE BEEN TERMINATED BEFORE GETTING THE RESULTS
// IN THAT CASE, A WAITGROUP SHOULD HAVE BEEN THE RIGHT CHOICE

package main

import (
	"fmt"
	"time"
)

func main() {

	requests := []string{"first", "second", "third"}
	responses := []int{}
	// I COULD HAVE USED AN UNBUFFERED (LENGTH=0) CHANNEL TOO.. THE BUFFERED ONE ALLOWS YOU TO FILL IT IN THE MAIN GOROUTINE WITH NO PROBLEMS, BUT HERE WE FILL THE CHANNEL ALWAYS INSIDE OTHER GOROUTINES
	respChan := make(chan int, len(requests))

	// perform requests asynchronously saving results in respChan
	for _, req := range requests {
		// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	        // IT'S IMPORTANT TO CREATE A COPY OF THE VARIABLE WHEN USING IN A GOROUTINE INSIDE A CYCLE: OTHERWISE TROUBLE!!
		// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	        request := req
	        // I COULD HAVE ALSO CREATED A FUNCTION PASSING THE CHANNEL AS PARAM INSTEAD OF USING THE CHANNEL INSIDE AN ANONYMOUS FUNCTION
		go func() {
			respChan <- httpCall(request)
		}()
	}

	// collecting responses with a blocking while cycle
	i := 0
	// IF YOU MISTAKE THE MAX NUMBER OF ITERATIONS IN EXCESS YOU GET A DEADLOCK ERROR BECAUSE THE MAIN GOROUTINE WAITS INDEFINITELY FOR ANOTHER VALUE TO POP UP IN THE CHANNEL: THING THAT WILL NEVER HAPPEN
	for i < len(requests) {
		// this operation blocks the goroutine until a new value is fetched
		responses = append(responses, <-respChan)
		i++
	}

	fmt.Println(responses)

}

func httpCall(requestBody string) (response int) {
	// simulate HTTP call
	fmt.Printf("performing the %s call...\n", requestBody)
	time.Sleep(6 * time.Second)
	fmt.Printf("... call %s executed!\n", requestBody)
	return 200
}

