package main

import "fmt"

func main() {
	a:= make([]int,0,10)
		fmt.Println(cap(a))
}

func s(s []int){
	s=append(s,0)
	for i := range s {
		s[i]++;
	}
}

func s3()  {
	s1 :=[]int{1,2}
	s2:=s1
	s2=append(s2,3)
	s(s1)
	s(s2)
	fmt.Println(s1,s2)
}