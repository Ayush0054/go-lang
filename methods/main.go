package main

import "fmt"

func main(){
	fmt.Println("methods in golang")
	hitesh := User{"Hitesh","hitesh@gad.com",true,16}
	// fmt.Printf("hitesh details are : %+v\n",hitesh)
	fmt.Println(hitesh.GetStatus())
	hitesh.NewMail()
	fmt.Printf("name is %v and email is  %v ",hitesh.Name,hitesh.Email)
}
type User struct{
	Name string
	Email string
	Status bool
	Age int
}
func (u User) GetStatus() string{
	if u.Status == true{
		return "active"
	}
	return "inactive"
}
func (u User) NewMail() {
	u.Email = "test@go.com"
	fmt.Println("Email of this user is ",u.Email)
}