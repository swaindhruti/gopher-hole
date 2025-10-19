package basics

import (
	"fmt"
	"math"
)

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func Methods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

/*
In Go, you can define methods on types. A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
In this example, we define a method Abs on the type Vertex. The method calculates the absolute value (magnitude) of the vector represented by the Vertex.
To call a method, you use the dot notation: v.Abs() calls the Abs method on the Vertex instance v.
*/
