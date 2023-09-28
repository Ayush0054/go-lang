package main

import "fmt"

func main(){
	fmt.Println("maps in golang")
  languages := make(map[string]string)
  languages["Js"]="javsccript"
  languages["Py"]="python"
  languages["rs"]="rust"
  fmt.Println("list of all languages: ",languages)
  fmt.Println("js shorts for: ",languages["Js"])
  delete(languages,"rs")
  fmt.Println("list of all languages: ",languages)

  for key , value := range languages{
	fmt.Printf("for key %v , value is %v\n",key,value)
  }
}