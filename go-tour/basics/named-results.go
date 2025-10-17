package basics

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/*
Go's return value can be named. If so, they are treated as variables defined at the top of the function.
A return statement without arguments returns the named return values. This is known as a "naked" return.
Naked returns should be used only in short functions, as they can harm readability.
*/
