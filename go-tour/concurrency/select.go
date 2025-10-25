package concurrency

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Select() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for range 10 {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

/*
The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case.
It chooses one at random if multiple are ready.

In this example, the `fibonacci` function generates Fibonacci numbers and sends them to channel `c`.
It also listens on the `quit` channel to know when to stop generating numbers.

The `Select` function starts a goroutine that receives 10 Fibonacci numbers from channel `c`
and then sends a signal on the `quit` channel to stop the Fibonacci generation.

This pattern is useful for managing multiple channel operations and coordinating goroutines.
*/
