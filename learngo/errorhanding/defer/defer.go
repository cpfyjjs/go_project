package main

import "fmt"

func tryDefer(){
	defer fmt.Println(1)
	defer fmt.Println(2)
	//panic("error occurred")
	fmt.Println(3)
}



func main(){
	defer fmt.Println(4)
	tryDefer()
}
