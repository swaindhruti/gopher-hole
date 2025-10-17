package basics

import "fmt"

func Defer() {
	fmt.Println("counting")

	for i := range 3 {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

/*
A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
Deferred functions are executed in last-in-first-out order after the surrounding function returns.
*/
