package main

import "fmt"

func main() {
	for row := 1; row <= 5; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
