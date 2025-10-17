package basics

import "fmt"

func For() {
	sum := 0
	for i := range 10 {
		sum += i
	}
	fmt.Println("Sum:", sum)
}
