package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printFileContents(f)
	file, err := os.OpenFile("fib.txt", os.O_APPEND|os.O_CREATE, 0666)
	if pathError, ok := err.(*os.PathError); ok {
		fmt.Println(pathError.Err, pathError.Op, pathError.Path)
	} else {
		fmt.Println(err)
	}
	file.Write([]byte("hello"))
}
