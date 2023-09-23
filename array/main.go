package main

import "fmt"

func main() {
	println("welcome to arrays in golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	// fruitList[2] = "Orange"
	fruitList[3] = "Peach"
   fmt.Println("Fruit list is ", fruitList)  
   fmt.Println("Fruit list is ", len(fruitList))

   var vegList = [3]string{"potato", "beans", "brinjal"}
   fmt.Println("veg list is ", vegList)
   fmt.Println("veg list is ", len(vegList))
}