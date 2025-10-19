package basics

import "fmt"

func Pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(p)  // print the pointer address
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

/*
The & operator generates a pointer to its operand. This is calles referencing.
The * operator denotes the pointer's underlying value. This is called dereferencing.
If p is a pointer, then *p is the value pointed to by p.
Deferencing a nil pointer will cause a runtime panic.
*/
