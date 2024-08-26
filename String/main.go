package main

import "fmt"

func main() {

	age := 30
	name := "Alex"
	value := 1000.476
	bul := true
	fmt.Println("My age is " + fmt.Sprint(age)) //My age is 30
	fmt.Printf("My age is %d", age)             //My age is 30				%d used for int
	fmt.Printf("\nMy name is %s", name)         //My name is Alex           %s for string
	fmt.Printf("\nI have %f", value)            //I have 1000.476000        %f for flouat
	fmt.Printf("\n%t", bul)                     //true         				%t for bool

}
