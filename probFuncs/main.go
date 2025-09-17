package main

import "fmt"

func add(a, b int) int {
	sum := a + b
	return sum
}

func area(length, width int) int {
	return length * width
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

var square = func(x int, y int) {
	n := x * y
	fmt.Println(n)
}

func operate(a, b int, op func(int, int) int) int {
	opr := op(a, b)
	return opr
}

func main() {
	add(1, 1)
	add(2, 2)
	rect := area(23, 20)
	fmt.Println(rect)
	even := isEven(1)
	fmt.Println(even)
	square(2, 5)
	dd := operate(3, 4, add)
	fmt.Println(dd)

}

func init() {
	func(x int, y int) {
		total := x * y
		fmt.Println(total)
	}(6, 7)
}
