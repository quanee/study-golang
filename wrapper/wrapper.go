package main

import "fmt"

type FibI interface {
	Fib(n int) int
	Wrap(fib FibI) FibI
}

type Fib struct {
	Wrapper FibI
}

func (this *Fib) Fib(n int) int {
	//wrapper := this.Wrapper
	if this.Wrapper == nil {
		this.Wrapper = this
	}
	fmt.Printf("Fib.Fib..%T...%v\n", this.Wrapper, n)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// call son
	fmt.Printf("Fib.Fib..%T...\n", this.Wrapper)
	return this.Wrapper.Fib(n-1) + this.Wrapper.Fib(n-2)
}

func (this *Fib) Wrap(fib FibI) FibI {
	fmt.Printf("Fib.Wrap..%T...\n", this.Wrapper)
	this.Wrapper = fib
	fmt.Printf("Fib.Wrap..%T...\n", this.Wrapper)
	fmt.Println("Fib.Wrap 调用", this)
	return this
}

type CacheFib struct {
	Wrapper FibI
	cache   map[int]int
}

func (this *CacheFib) Wrap(fib FibI) FibI {
	this.Wrapper = fib
	fmt.Printf("CacheFib.Wrap..%T...\n", this.Wrapper)
	return this
}

func (this *CacheFib) Fib(n int) int {
	if this.cache == nil {
		this.cache = make(map[int]int)
	}
	if ans, ok := this.cache[n]; ok {
		return ans
	}
	ans := this.Wrapper.Fib(n)
	this.cache[n] = ans
	return ans
}

type CounterFib struct {
	Wrapper FibI
	Counter int
}

func (this *CounterFib) Wrap(fib FibI) FibI {
	this.Wrapper = fib
	fmt.Printf("CounterFib.Wrap..%T...\n", this.Wrapper)
	return this
}

func (this *CounterFib) Fib(n int) int {
	this.Counter++
	return this.Wrapper.Fib(n)
}

func main() {
	fib := new(Fib)
	//fmt.Println("result fib", fib.Fib(10))

	cacheFib := new(CacheFib)
	//f := cacheFib.Wrap(fib.Wrap(cacheFib))
	//fmt.Println(f.Fib(10), "heihei")
	counterFib := new(CounterFib)
	counterCacheFib := cacheFib.Wrap(counterFib.Wrap(fib.Wrap(cacheFib)))
	fmt.Println("result cache:counter:fib", counterCacheFib.Fib(10))
	fmt.Printf("count: %d, cache: %v", counterFib.Counter, cacheFib.cache)
}
