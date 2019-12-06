package main

func f1(p *int) int {
	*p = 123
	return *p
}

func g(x int) (a, b int) {
	return x, f1(&x)
}

/*func main() {
	fmt.Println(g(0))
}*/
