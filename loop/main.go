package main

import "fmt"

func main(){
 fmt.Println("loops in golang")
 days := []string{"sun","tues","wed","thu","fri","sat"}
 fmt.Println(days)
  //first type
 //  for d:=0; d< len(days); d++{
// 	fmt.Println(days[d])
//  }

//second type
// for i:= range days{
// 	fmt.Println(days[i])
// }

//third type  (for eacch type)

// for index, day := range days{
// 	fmt.Printf("index is %v and value is %v\n",index,day)
// }
//-------------------------------------------------------------
rouguevalue := 1 
for rouguevalue < 10 {
	if rouguevalue == 2{
		goto lco
	}
	if rouguevalue ==5{
		break
	}
	fmt.Println("value is ",rouguevalue)
	rouguevalue++;
}

lco: 
fmt.Println("juumping in lco")
}