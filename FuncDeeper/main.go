package main

import "fmt"

////func fn(someFunctuin func(int) int) {
////	fmt.Println(someFunctuin(25))
////}

func test(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {

	// a := func(x int, y int) int {
	// 	return x + y
	// }
	// sum := a(10, 4)
	// fmt.Println(sum)

	//// square := func(x int) int {
	//// 	return x * x
	//// }

	//// fn(square)
	// a := test("Hello") OR
	test("Hello")()
}
