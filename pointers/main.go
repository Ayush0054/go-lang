package main

import (
	"fmt"
)

func main(){
	fmt.Println("welcome to pointers in golang")
	// var ptr *int 
	// fmt.Println("value of pointer is ", ptr)
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("value of pointer is ", ptr)
	fmt.Println("value of pointer is ", *ptr)
	*ptr = *ptr + 1
	fmt.Println("value of pointer is ", *ptr)
}