package basics

func TypeConversions() (int, float64, byte) {
	var i int = 42
	var f float64 = float64(i)
	var b byte = byte(f)

	return i, f, b
}

/*
Go is a statically typed language, which means that every variable has a type that is known at compile time.
However, Go also provides a way to convert between types. To convert a value of one type to another type, you use the syntax T(v), where T is the target type and v is the value you want to convert.

In this example, we convert an integer to a float64 and then to a byte. Note that converting from a larger type to a smaller type (like from int to byte) can result in data loss if the value is out of range for the target type.
*/
