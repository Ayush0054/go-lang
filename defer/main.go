package main

import "fmt"

func main(){
	// defer  fmt.Println("defer in golang")
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")
	myDefer()
}
func myDefer(){
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}

}