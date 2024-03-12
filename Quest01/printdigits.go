package main

import "fmt"

func PrintDigits_Ascending() {
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

func PrintDigits_Descending() {
	for i := 9; i >= 0; i-- {
		fmt.Printf("%d", i)
	}
}
func main() {
	PrintDigits_Ascending()
	PrintDigits_Descending()
}
