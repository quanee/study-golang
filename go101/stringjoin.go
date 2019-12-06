package main

var s string
var xa = []byte{1023: 'x'}
var ya = []byte{1023: 'y'}

func fc() {
	s = (" " + string(xa) + string(ya))[1:]
}
func fd() {
	s = string(xa) + string(ya)
}

//func main() {
//	fc()
//	fmt.Println(s)
//}
