package main

import "fmt"

/*import "GoStudy/cgo"

//
int func main() {
	cgo.CHelloworld()
}*/

func aa(a * int)  {
	*a = *a + 1
	return
}

func aaa(a int) int {
a=a+1
return a
}

func main() {

	a:=10;
aaa(a)
	fmt.Println(a)

	aa(&a);
	fmt.Println("aa:",a)
}