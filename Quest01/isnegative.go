package main

import "fmt"

func Isnegative(nbr int){
	if nbr < 0 {
		fmt.Println("T")
	}else{
		fmt.Println("F")
	}
}
func main(){
	Isnegative(5)
	Isnegative(0)
	Isnegative(-5)
}