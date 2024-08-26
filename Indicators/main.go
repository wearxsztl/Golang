package main

import "fmt"

func change(str *string) {
	// *str = "Lol"
}

func main() {
	// a := 5
	// pointer := &a
	// fmt.Println(pointer)  //0xc00000a0d8
	// fmt.Println(*pointer) //5 * оператор разыменования
	// s := "uuu"
	// var pointer *string = &s
	// fmt.Println(s)
	// change(&s)
	// fmt.Println(s)
	num := 9
	b := &num
	*b = 35
	fmt.Println(*b)

}
