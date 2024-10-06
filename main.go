package main

// #cgo CFLAGS: -I/opt/homebrew/Cellar/python@3.12/3.12.7/Frameworks/Python.framework/Versions/3.12/include/python3.12
// #cgo LDFLAGS: -L/opt/homebrew/Cellar/python@3.12/3.12.7/Frameworks/Python.framework/Versions/3.12/lib -lpython3.12
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
