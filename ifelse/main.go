package main

import "fmt"

func main(){
	fmt.Println("if else in golang")
	loginCount := 6
	var result string
	 if loginCount < 10 {
       result = " regular user"
	 }else if loginCount > 10{
		result = " watch out"
	 }else{
		result ="exactly 10 login count"
	 }
	 fmt.Println(result)

    if 9%2 == 0{
		fmt.Println("even")
	}else{
		fmt.Println("odd")
	}
 if num := 3; num<10{
	fmt.Println("num is les than 10")
 }else{
	fmt.Println("num more then 10")
 }
	}