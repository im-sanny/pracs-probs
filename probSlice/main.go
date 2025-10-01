package main

import "fmt"

// func changeSlice(p []int) []int {
// 	p[0] = 10         // [10, 6, 7] overwrite 5 from x
// 	p = append(p, 11) // [10, 6, 7, 11] new array creation -> disconnect y from  a and x
// 	return p          // return for y
// }

func main() {
	// x := []int{1, 2, 3, 4, 5}
	// x = append(x, 6) // [1, 2, 3, 4, 5, 6]
	// x = append(x, 7) // [1, 2, 3, 4, 5, 6, 7]

	// a := x[4:]          // [5, 6, 7]
	// y := changeSlice(a) //sending a in changeSlice. got [10, 6, 7, 11] from changeSlice

	// fmt.Println(x) //[1, 2, 3, 4, 10, 6, 7]
	// fmt.Println(y) //[10, 6, 7, 11]

	// a := []int{1, 2, 3}
	// b := a //[1, 2, 3]
	// a = append(a, 4) // [1, 2, 3, 4]
	// b[0] = 99 // [99, 2, 3]
	// fmt.Println(a)
	// fmt.Println(b)

	x := make([]int, 2, 3)
	x[0], x[1] = 1, 2 // [1, 2]
	y := x            // [1, 2]
	x = append(x, 3)  // [1, 2, 3]
	y = append(y, 4)  // [1, 2, 4] // after allocation y gets disconnected from x
	x[0] = 99         //[99, 2, 4]
	fmt.Println(x)    //[99, 2, 4]
	fmt.Println(y)    // [1, 2, 4]

	p := []int{1, 2}
	q := p
	p = append(p, 3) //[1, 2, 3]
	q = append(q, 4) //[1, 2, 3,4]
	p[1] = 88        //[2,88, 3]
	q[0] = 77        //[77, 2, 3, 4]
	fmt.Println(p)   //
	fmt.Println(q)   //

	// I guess my mistake in these problem is I'm not calculating len and cap

	m := []int{1, 2, 3, 4}
	n := m[:2]       // [1, 2]
	m = append(m, 5) //[1, 2, 3, 4, 5]
	n = append(n, 6) //[1, 2, 6]
	m[0] = 99        //[99, 2, 3, 4, 5]
	fmt.Println(m)   //[99, 2, 3, 4, 5]
	fmt.Println(n)   //[1, 2, 6]

	s1 := []int{1, 2}
	s2 := s1
	s1 = append(s1, 3) // [1,2,3]
	s2 = append(s2, 4) // [1, 2, 4]
	s1 = append(s1, 5) // [1,2,3,5]
	s2[0] = 99         // [99, 2, 4]
	fmt.Println(s1)    // [1,2,3,5]
	fmt.Println(s2)    // [99, 2, 4]

}
