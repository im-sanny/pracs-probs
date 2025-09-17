package main

import "fmt"

var a = 1

func add(x int, y int) {
	sum := x + y
	fmt.Println(sum)
}

func totalSum(c int, d int, e int) int {
	sum := c + d + e
	return sum
}

func shadow() {
	a := 5
	fmt.Println(a + 5)
}

func main() {
	b := 2
	g := 3

	add(a, b)
	fmt.Println(totalSum(a, b, g))
	shadow()
}
