package main

/*
#include <errno.h>
void myfunc() { errno = EINVAL; }
int myfunc2() { errno = EINVAL; return 8; }
*/
import "C"

import "fmt"

//import "syscall"

func main() {
	_, err := C.myfunc()     // assigns void instead of syscall.Errno
	fmt.Printf("%#v\n", err) // main._Ctype_void{}

	//var i C.int
	//var err2 error
	//i, err2 := C.myfunc2()                     // works as expected
	var i C.int
	var err2 error
	i, err2 = C.myfunc2()                     // works as expected
	fmt.Printf("%#v %#v %T\n", i, err2, err2) // 8 0x16 syscall.Errno
	// fmt.Printf("%#v\n", err2.(syscall.Errno)) // syscall.Errno
}
