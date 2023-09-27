package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("welcome to slices in golang")

	var fruitList = []string{"Apple", "Tomato", "Orange"}
	fmt.Printf("Fruit list is %T\n ", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fruitList = append(fruitList[1:])
	fmt.Println("Fruit list is ", fruitList)
	highScores := make([]int, 4)
	highScores[0] = 234		
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4]=666 //wont work
    highScores = append(highScores, 555,2342)  //will work
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	var courses = []string{"reactjs","vue","angular","node"}
	fmt.Println(courses)
	
}