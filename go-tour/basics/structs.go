package basics

import "fmt"

type Vertex struct {
	X, Y int
}

func Structs() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v.X)

	p := &v   // point to v
	p.Y = 1e9 // set Y through the pointer
	fmt.Println(v)
}

/*
Struct fields can be accessed through a struct pointer.
For a struct pointer p, the expression (*p).X is equivalent to p.X.
The special notation p.X is shorthand for (*p).X.
*/
