package main

import "C"
import "fmt"

//export GetIntFromDLL
func GetIntFromDLL() int32 {
  return 42
}

//export PrintHello
func PrintHello(name *C.char) {
  fmt.Printf("From DLL: Hello, %s!\n", C.GoString(name))
}

//export PrintBye
func PrintBye() {
  fmt.Println("From DLL: Bye!")
}

func main() {
  // dummy function to make CGO compile package as C shared library
}
