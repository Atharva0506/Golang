package main

/*
#include <stdio.h>
#include <stdlib.h>

void my_c_print(char* str) {
    printf("Hello from the C Language: %s\n", str);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("This is printed by Go!")

	// We must convert the Go string into a C-compatible String!
	cString := C.CString("My name is Gopher!")

	// IMPORTANT: C does not have an automatic Garbage Collector!
	// We MUST manually free the memory when we are done!
	defer C.free(unsafe.Pointer(cString))

	// Call the C function perfectly natively!
	C.my_c_print(cString)
}
