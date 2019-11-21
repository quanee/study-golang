package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"study-golang/pipline"
)

func main() {
	//p := createPipeline("large.in", 80000000, 4)
	p := createNetworkPipeline("large.in", 80000000, 4)
	//p := createNetworkPipeline("small.in", 512, 4)
	writeToFile(p, "large.out")
	printFile("large.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipline.WriterSink(writer, p)
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipline.Init()
	sortResults := make([]<-chan int, 0)
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipline.InMemSort(source))
	}

	return pipline.MergeN(sortResults...)
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipline.Init()
	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(700 + i)
		pipline.NetworkSink(addr, pipline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipline.NetworkSource(addr))
	}
	return pipline.MergeN(sortResults...)
}
