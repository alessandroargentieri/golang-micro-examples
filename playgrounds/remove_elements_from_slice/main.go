package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a slice of time.Time values
    times := []time.Time{
        time.Now(),                   // Current time
        time.Now().Add(-20 * time.Minute), // 20 minutes ago
        time.Now().Add(-5 * time.Minute),  // 5 minutes ago
        time.Now().Add(-30 * time.Minute), // 30 minutes ago
    }

    // Loop over the slice in reverse order
    for i := len(times) - 1; i >= 0; i-- {
        elapsed := time.Since(times[i])
        if elapsed.Minutes() > 10 {
            // If more than 10 minutes have passed, remove the element from the slice
            times = append(times[:i], times[i+1:]...)
        }
    }

    fmt.Println(times) // Output: [current time 5 minutes ago]
}
/*
In this example, we loop over the times slice in reverse order using a for loop with an index variable i. We use the 
time.Since() function to calculate the elapsed time for each time.Time value in the slice. If the elapsed time is greater 
than 10 minutes, we use the append() function to remove the element from the slice by creating a new slice that excludes 
that element.

Note that when using the append() function to remove an element from a slice, we need to create a new slice by concatenating 
the slice before the element (times[:i]) and the slice after the element (times[i+1:]). This ensures that we don't access 
the slice out of bounds or cause any other slice access issues.

Why do we iterate in a reverse order?

In the example provided, we iterate over the slice in reverse order because we're removing elements from the slice while 
iterating.

When you remove an element from a slice, the slice shrinks and the indexes of all subsequent elements shift down by one. 
This can cause problems if you're iterating over the slice from beginning to end, as removing an element will change the 
indexes of the remaining elements, and you might end up skipping an element or processing an element twice.

By iterating over the slice in reverse order, we avoid this issue. Since we're removing elements from the end of the slice 
first, the indexes of the remaining elements don't change, and we don't skip any elements or process any element twice.

It's worth noting that iterating over a slice in reverse order does add some additional overhead compared to iterating over 
it from beginning to end. However, in cases where you need to remove elements from the slice while iterating, it's usually the safer and more reliable approach.
