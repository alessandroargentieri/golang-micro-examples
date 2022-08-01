package main

import (
	"fmt"
)

func main() {
	
	// BREAK AN INFINITE LOOP ****************
	i := 1
	for {
	   i++
	   if i == 5 {
	     break;
	   }
	}
	fmt.Println(i)
	

	// WHILE LOOP ****************************
	i = 1
	for i <= 4 {
	   i++
	}
	fmt.Println(i)
	
	
	// FOR LOOP ******************************
	for i=1; i <=5; i++ {
	   if i == 5 {
	      fmt.Println(i)
	   }
	}
	
	
	// RANGE LOOP ON SLICE *******************
	for i, v := range []int{ 1, 2, 3, 4, 5 } {
	    fmt.Println(i, " ", v)
	}
	
	// RANGE LOOP ON MAP *********************
	for k, v := range map[int]int{ 1:1, 2:2, 3:3, 4:4, 5:5 } {
	    fmt.Println(k, " ", v)
	}
	
	// FOR LOOP ON CHAN **********************
	mychan := make(chan int, 5)
	mychan <- 1; mychan <- 2; mychan <- 3; mychan <- 4; mychan <- 5;
	close(mychan) 
	for v := range mychan {
	    fmt.Println(v)
	}
	
}

