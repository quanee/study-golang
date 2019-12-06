package main

var (
	_  = f2("w", x1)
	x1 = f2("x1", z)
	y  = f2("y", x1)
	z  = f2("z")
)

func f2(s string, deps ...int) int {
	print(s)
	return 0
}

/*func main() {

}*/
