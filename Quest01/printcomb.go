package main

import "fmt"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				fmt.Printf("%d", i)
				fmt.Printf("%d", j)
				fmt.Printf("%d", k)

				if i != 7 || j != 78 || k != 9 {
					fmt.Print(",")
					fmt.Print("")
				}

			}
		}
	}
}

func main() {
	PrintComb()
}
