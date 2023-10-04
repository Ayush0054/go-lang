package main

import (
	"fmt"
	"net/url"
)
const myurl string  = "https://www.youtube.com/watch?v=lG7Uxts9SXs"
func main(){
	fmt.Println("hello world")
	fmt.Println(myurl)
	//parsing 

	result, _ :=url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	// fmt.Println(result.Query())
	// fmt.Println(result.Fragment)
	// fmt.Println(result.User)
	// fmt.Println(result.Opaque)

	qparams := result.Query()
	fmt.Printf(" the type is : %T\n",qparams)
	fmt.Println(qparams["v"][0])

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}
	fmt.Println(partsOfUrl.String())
	
}