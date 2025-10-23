package concurrency

import "fmt"

func Sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func Channels() {
	fmt.Println("We will learn about channels here.")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go Sum(s[:len(s)/2], c)
	go Sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println("Sum:", x, y, "Total:", x+y)
}

/*
Working of this example:

1. A unbuffered channel `c` of type `int` is created using `make(chan int)`.
2. Two goroutines are started using the `go` keyword, each executing the `Sum` function on different halves of the slice `s`.
3. Each `Sum` function calculates the sum of its respective slice and sends the result to the channel `c` using `c <- total`.
4. The main function waits to receive two values from the channel `c` using `<-c`, which blocks until a value is available.
5. Finally, it prints the individual sums and their total.

Channels in Go provide a way for goroutines to communicate with each other and synchronize their execution.
*/
