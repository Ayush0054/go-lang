package main

import "fmt"

const LoginToken string ="fasddsf"
func main() {
	var username string = "ayush"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)
	var smallFLoat float32  = 4214.4132444444444444444444444444444444444
	fmt.Println(smallFLoat)
	fmt.Printf("variable is of type: %T \n", smallFLoat)

	//default values and some aliases
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	//implicit type 
	var website = "ayush.so"
	fmt.Println(website)

	// no var style
	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)
    fmt.Println(LoginToken)
}