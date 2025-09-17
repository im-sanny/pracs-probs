package main

import "fmt"

func add(a int, b int) { // a and b is parameter here
	fmt.Println(a + b)
}

func hof(x int, y int, op func(c int, d int)) {
	op(x, y) //callback
}

func mult(x int, y int)func() int{
	return func () int {
		return x * y
	}
}


func main() {
	res := mult(2, 2)
	fmt.Println(res())
	hof(1, 2, add) // this is argument
	// fmt.Println(mult(2, 2)())
}

func init() {
	func(a string) {
		fmt.Println(a)
	}("IIFE and Anonymous")
}

/*
From main function I called hof with 3 params, 2 of them int and one is func.
When the main func is called, first it’ll send 1 and 2 into hof, then x = 1 and y = 2.
Inside hof, op is the function parameter. Since I passed add from main, now op = add.

When I call op(x, y), it’s actually like calling add(x, y). So x and y (1 and 2) are passed into add, and add prints the result.
That’s why I don’t need to return op or print inside hof, because add already handles the printing.
*/

/*
1. parameter and argument
2. standard func
3. HOF
4. callback
5. init
6. IIFE
7. anonymous
*/
