package basics

import "fmt"

func (v *Vertex) Scale(f float64) {
	v.X = int(float64(v.X) * f)
	v.Y = int(float64(v.Y) * f)
}

func PointersMethod() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v)
}

/*
Methods can be defined for pointer receivers.
In this example, we define a method Scale that scales the X and Y fields of the Vertex by a given factor f.
The receiver is a pointer to Vertex (*Vertex), which allows the method to modify the original Vertex value.
When we call v.Scale(10), Go automatically takes the address of v to match the pointer receiver type.
*/
