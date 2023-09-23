package main

import (
	"bufio"
	"fmt"
	"os"
)
func main (){
	welcome := "Welcome to my playground"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || err err syntax
	input , _ :=reader.ReadString('\n')
	fmt.Println("thanks for rating,",input)
	fmt.Printf("type of this  rating %T,",input)
}