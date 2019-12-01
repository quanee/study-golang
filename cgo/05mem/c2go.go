package main

/*
extern char* NewGoString(char* );
extern void FreeGoString(char* );
extern void PrintGoString(char* );

static void printString(char* s) {
	char* gs = NewGoString(s);
	PrintGoString(gs);
	FreeGoString(gs);
}
*/
import "C"
import (
	"sync"
	"unsafe"
)

type ObjectId int32

var refs struct {
	sync.Mutex
	objs map[ObjectId]interface{}
	next ObjectId
}

func init() {
	refs.Lock()
	defer refs.Unlock()

	refs.objs = make(map[ObjectId]interface{})
	refs.next = 1000
}

func NewObjectId(obj interface{}) ObjectId {
	refs.Lock()
	defer refs.Unlock()

	id := refs.next
	refs.next++

	refs.objs[id] = obj
	return id
}

func (id ObjectId) IsNil() bool {
	return id == 0
}

func (id ObjectId) Get() interface{} {
	refs.Lock()
	defer refs.Unlock()

	return refs.objs[id]
}

func (id *ObjectId) Free() interface{} {
	refs.Lock()
	defer refs.Unlock()

	obj := refs.objs[*id]
	delete(refs.objs, *id)
	*id = 0

	return obj
}

//export NewGoString
func NewGoString(s *C.char) *C.char {
	gs := C.GoString(s)
	id := NewObjectId(gs)
	return (*C.char)(unsafe.Pointer(uintptr(id)))
}

//export FreeGoString
func FreeGoString(p *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(p)))
	id.Free()
}

//export PrintGoString
func PrintGoString(s *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(s)))
	gs := id.Get().(string)
	print(gs)
}

func main() {
	C.printString("hello")
}
