package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lmath
#include <math.h>
*/
import "C"

import "fmt"

func main() {
	result := C.add(1, 2)
	fmt.Println("Sum is ", result)
}
