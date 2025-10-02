package main

import "fmt"

func changeSlice(p []int) []int {
	p[0] = 10         // [10, 6, 7] overwrite 5 from x
	p = append(p, 11) // [10, 6, 7, 11] new array creation -> disconnect y from  a and x
	return p          // return for y
}

func prob1() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6) // [1, 2, 3, 4, 5, 6]
	x = append(x, 7) // [1, 2, 3, 4, 5, 6, 7]

	a := x[4:]          // [5, 6, 7]
	y := changeSlice(a) //sending a in changeSlice. got [10, 6, 7, 11] from changeSlice

	fmt.Println(x) //[1, 2, 3, 4, 10, 6, 7]
	fmt.Println(y) //[10, 6, 7, 11]
}

func prob2() {
	a := []int{1, 2, 3}
	b := a           // [1, 2, 3]
	a = append(a, 4) // [1, 2, 3, 4]
	b[0] = 99        // [99, 2, 3]
	fmt.Println(a)   // [1, 2, 3, 4]
	fmt.Println(b)   // [99, 2, 3]
}

func prob3() {
	x := make([]int, 2, 3) //len = 2, cap = 3
	x[0], x[1] = 1, 2      // [1, 2]
	y := x                 // [1, 2]
	x = append(x, 3)       // [1, 2, 3] len= 3, cap = 3
	y = append(y, 4)       // [1, 2, 4] len= 3 cap = 3, here 4 will overwrite 3, bc y had 2 len and 3 cap so it won't create new array and that's why it will change the original array.
	x[0] = 99              // [99, 2, 4] // since there's no new array and no disconnection so y will get this change of x.
	fmt.Println(x)         // [99, 2, 4]
	fmt.Println(y)         // [99, 2, 4]
}

func prob4() {
	p := []int{1, 2}
	q := p           // [1, 2]
	p = append(p, 3) //[1, 2, 3]
	q = append(q, 4) //[1, 2, 4] // After p = append(p, 3), p and q point to completely separate arrays. They are no longer connected.
	p[1] = 88        //[1, 88, 3]
	q[0] = 77        //[77, 2, 4]
	fmt.Println(p)   //[1, 88, 3]
	fmt.Println(q)   //[77, 2, 4]
}

func prob5() {
	m := []int{1, 2, 3, 4}
	n := m[:2]       // [1, 2]
	m = append(m, 5) //[1, 2, 3, 4, 5]
	n = append(n, 6) //[1, 2, 6]
	m[0] = 99        //[99, 2, 3, 4, 5]
	fmt.Println(m)   //[99, 2, 3, 4, 5]
	fmt.Println(n)   //[1, 2, 6]
}

func prob6() {
	s1 := []int{1, 2}
	s2 := s1
	s1 = append(s1, 3) // [1,2,3]
	s2 = append(s2, 4) // [1, 2, 4]
	s1 = append(s1, 5) // [1,2,3,5]
	s2[0] = 99         // [99, 2, 4]
	fmt.Println(s1)    // [1,2,3,5]
	fmt.Println(s2)    // [99, 2, 4]
}

func prob7() {
	a := []int{1, 2, 3}
	b := a[:2]       // [1, 2]
	a = append(a, 4) // [1, 2, 3, 4]
	b = append(b, 5) // [1, 2, 5]
	a[0] = 99        // [99, 2, 3, 4]
	fmt.Println(a)   // [99, 2, 3, 4]
	fmt.Println(b)   // [1, 2, 5]
}

func prob8() {
	x := make([]int, 3, 5) //len = 3, cap = 5
	x[0], x[1], x[2] = 1, 2, 3
	y := x[:4]       // [1, 2, 3, 0]
	x = append(x, 4) // Uses spare capacity, no new array
	y[3] = 88        // Overwrites the 4 with 88
	fmt.Println(x)   //[1, 2, 3, 88]
	fmt.Println(y)   //[1, 2, 3, 88]
}

func prob9() {
	p := []int{1, 2}
	q := p
	r := q
	p = append(p, 3) // [1, 2, 3]
	q = append(q, 4) // [1, 2, 4]
	r[0] = 99        // [99, 2]
	fmt.Println(p)   // [1, 2, 3]
	fmt.Println(q)   // [1, 2, 4]
	fmt.Println(r)   // [99, 2]
}

func prob10() {
	m := []int{1, 2, 3}
	n := m[1:3]      // [2, 3]
	m = append(m, 4) // creates a new array for reallocation so it won't get any changes made in other array that share same array
	n[0] = 88
	n = append(n, 5) //[88, 3, 5]
	fmt.Println(m)   //[1, 2, 3, 4]
	fmt.Println(n)   //[88, 3, 5]
}

func prob11() {
	s1 := []int{1, 2}
	s2 := s1
	s1 = append(s1, 3, 4) // adding 2 elements
	s2 = append(s2, 5)
	s1[0] = 99
	fmt.Println(s1) //[99 2 3 4]
	fmt.Println(s2) //[1 2 5]
}

func main() {
	prob1()
	prob2()
	prob3()
	prob4()
	prob5()
	prob6()
	prob7()
	prob8()
	prob9()
	prob10()
	prob11()

}
