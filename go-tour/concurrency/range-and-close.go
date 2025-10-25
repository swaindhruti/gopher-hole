package concurrency

import "fmt"

func RangeAndClose() {
	fmt.Println("We will learn about range and close with channels here.")

	c := make(chan int)

	// Start a goroutine to send values to the channel
	go func() {
		for i := range 5 {
			c <- i
		}
		close(c) // Close the channel after sending all values
	}()

	// Use range to receive values from the channel until it's closed
	for v := range c {
		fmt.Println(v)
	}
}

/*
Working of this example:

1. A channel `c` of type `int` is created using `make(chan int)`.
2. A goroutine is started that sends integers from 0 to 4 into the channel `c`.
3. After sending all values, the channel is closed using `close(c)`.
4. In the main function, a `for range` loop is used to receive values from the channel `c`.
   The loop continues until the channel is closed, at which point it exits gracefully.

Using `range` with channels allows for easy and safe iteration over values sent to the channel,
and closing the channel signals that no more values will be sent.
*/
