package pipline

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestInMemSort(t *testing.T) {
	p := Merge(
		InMemSort(ArraySource(3, 2, 6, 7, 4)),
		InMemSort(ArraySource(7, 4, 0, 3, 2, 8, 13)),
	)

	for v := range p {
		fmt.Println(v)
	}
}

func TestMerge(t *testing.T) {
	const filename = "small.in"
	const n = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := RandomSource(n)
	writer := bufio.NewWriter(file)
	WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	p = ReaderSource(reader, -1)
	i := 0
	for v := range p {
		fmt.Println(v)
		i++
		if i == 100 {
			break
		}
	}
}
