package main

import "fmt"

func pattern() {
	for row := 1; row <= 5; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func reverse() {
	s := "hello"
	rev := ""

	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	fmt.Println(rev)
}

func main() {
	pattern()
	reverse()

}
