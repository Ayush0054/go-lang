package main

import "fmt"

func main (){
	fmt.Println("functions in golang")
greeter()
 result := adder(3,5)
 fmt.Println("result is ", result)

 proResult , proMessage := proAdder(1,2,3,4,5,6,7,8,9)
 fmt.Println("proResult is ", proResult)
 fmt.Println("proMessage is ", proMessage)
}
func adder(valone int , valtwo int) int{
	return valone + valtwo
}
func proAdder(values ...int) (int, string) {
       total := 0
	   for _, val := range values{
		   total += val
	   }
	   return total ,"total"
}	
func greeter (){
	fmt.Println("yooo")
}