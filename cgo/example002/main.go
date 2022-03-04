package main

/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"

import "fmt"

func main() {
	res, err := C.sqrt(1) // works
	//var res, err = C.sqrt(1) // assignment count mismatch: 2 = 1
	fmt.Println(res, err)
}
