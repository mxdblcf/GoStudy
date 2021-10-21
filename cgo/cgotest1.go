package cgo

/*
    #include <stdio.h>
    #include <stdlib.h>
    void hello (){
    printf("Hello,world!\n");
    }
 */
import "C"
import (
	"fmt"
)

func CHelloworld() {
	cs := C.hello();
	fmt.Println(cs)
}
