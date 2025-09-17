package main

import "fmt"

// Basic operations
func add(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func area(length, width int) int {
	return length * width
}

func isEven(n int) bool {
	return n%2 == 0
}

// Anonymous function assigned to a variable
var square = func(x, y int) int {
	return x * y
}

// Higher-order function
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Closure example
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Runs before main
func init() {
	// Immediately invoked function
	func(x, y int) {
		fmt.Println("Init multiply:", x*y)
	}(6, 7)
}

func main() {
	// Area
	rect := area(23, 20)
	fmt.Println("Rectangle area:", rect)

	// Even check
	fmt.Println("Is that a even?", isEven(1))

	// Square (anonymous function)
	fmt.Println("Square:", square(2, 5))

	// Using operate with add and mult
	fmt.Println("Operate add:", operate(3, 4, add))
	fmt.Println("Operate mult:", operate(50, 5, mult))

	// Closure example
	addFive := makeAdder(5)
	fmt.Println("AddFive(10):", addFive(10))
}
