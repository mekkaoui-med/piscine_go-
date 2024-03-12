package main

import "fmt"

func PrintAlpha_Ascending(){
	for i:='a';i<='z';i++{
		fmt.Printf("%c",i)
	}
	fmt.Println()
}

func PrintAlpha_Descending(){
	for i:='z';i>='a';i--{
		fmt.Printf("%c",i)
	}
	fmt.Println()
}

func main(){
	PrintAlpha_Ascending()
	PrintAlpha_Descending()
	
}