package main

import "fmt"

func main(){
	fmt.Println("structs in golang")
	hitesh := User{"Hitesh","hitesh@gad.com",true,16}
	fmt.Printf("hitesh details are : %+v\n",hitesh)
	fmt.Printf("name is %v and email is  %v ",hitesh.Name,hitesh.Email)
}
type User struct{
	Name string
	Email string
	Status bool
	Age int
}