package main

// #cgo pkg-config: python3-embed
// #include <Python.h>
import "C"
import "fmt"

func main() {
	defer C.Py_Finalize()
	C.Py_Initialize()
	pyCodeStr := `print("Hello, World!")`
	pyCodeChar := C.CString(pyCodeStr)
	C.PyRun_SimpleString(pyCodeChar)
	fmt.Println(C.GoString(C.Py_GetVersion()))
}
