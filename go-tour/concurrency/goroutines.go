package concurrency

import (
	"fmt"
	"time"
)

func Say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Goroutines() {
	fmt.Println("This is a placeholder for the goroutines example.")
	go Say("Hello from a goroutine!")
	Say("Hola!")
}

/*
Goroutines are lightweight threads managed by the Go runtime.
They are created using the `go` keyword followed by a function call.
When a goroutine is created, it runs concurrently with other goroutines in the same address space.

In this example, we define a function `Goroutines` that currently serves as a placeholder.
To demonstrate goroutines, you would typically call a function with the `go` keyword,
which would allow that function to execute concurrently with the main program flow.
*/
