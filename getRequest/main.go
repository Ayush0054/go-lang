package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
getRequest()
}

func getRequest(){
 const myurl string  = "http://localhost:8080/get"
 response,err := http.Get(myurl)
 if(err != nil){
	 panic(err)
 }
 defer response.Body.Close()
//  fmt.Println(response.Body)
//  fmt.Println(response.contentLength)
 fmt.Println(response.StatusCode)
 Content, _ := ioutil.ReadAll(response.Body)
 fmt.Println(string(Content))
}