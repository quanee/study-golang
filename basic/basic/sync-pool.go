package main

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func main() {
	buffer := bufferPool.Get().(*bytes.Buffer)
	// 使用buffer
	buffer.WriteString("")
	// 放回Pool中
	buffer.Reset()
	bufferPool.Put(buffer)
}
