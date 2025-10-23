package concurrency

import "fmt"

func UnbufferedChannels() {
	c := make(chan string, 3) // Create a buffered channel with capacity 3

	c <- "Message 1"
	c <- "Message 2"
	c <- "Message 3"

	// The following line would block because the channel is full
	// c <- "Message 4"

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	// Now we can send another message since we've read from the channel
	c <- "Message 4"
	fmt.Println(<-c)
}

/*
In this example, we create a buffered channel `c` with a capacity of 3 using `make(chan string, 3)`.
We then send three messages to the channel without blocking because the channel has enough capacity to hold them.
If we were to try to send a fourth message before reading any from the channel, it would block until space is available.

After sending the messages, we read from the channel three times, which frees up space in the channel.
We can then send another message without blocking.

Buffered channels allow for more flexible communication between goroutines by providing a buffer that can hold multiple values.
This can help improve performance and reduce blocking in certain scenarios.
*/
