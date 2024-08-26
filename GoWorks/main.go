package main

import (
	"fmt"
	"math"
)

func main() {

	//fmt.Println("Hello world")
	//var a string
	//fmt.Scan(&a)

	//var age int8 = 25
	//fmt.Println(age)

	//var number float64 = 267.6
	//fmt.Println(number)

	//age := 16
	//fmt.Println(age)

	//var name string = "Alexandr"
	// name := "Alexandr"
	// fmt.Println(name)
	// fmt.Println(len(name))

	// var name string
	// var age string
	// fmt.Println("What is your name? ")
	// fmt.Scan(&name)
	// fmt.Println("How old are you? ")
	// fmt.Scan(&age)
	// fmt.Println("Hello " + name + ", " + fmt.Sprint(age) + " years")

	// var h int8 = 8
	// var t int64 = int64(h)
	// fmt.Print(t)

	// var num int64 = -3
	// if num > 0 {
	// 	fmt.Println("Number > 0 ")
	// } else if num < 0 {
	// 	fmt.Println("Number < 0")
	// } else if num == 3 {
	// 	fmt.Println("Number is 3")
	// }
	/////////////////////////////////////////////////////// Поиск дискриминанта и корней
	var a float64
	var b float64
	var c float64
	fmt.Println("Реши квадратное уравнение ")
	fmt.Println("Введи а:")
	fmt.Scan(&a)
	fmt.Println("Введи b:")
	fmt.Scan(&b)
	fmt.Println("Введи c:")
	fmt.Scan(&c)

	D := (b * b) - (4 * a * c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Уравнение имеет 2 корня\nD= " + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("Ваше уравнение имеет 1 корень\n D = 0 ")
		fmt.Println("X: " + fmt.Sprint(x))

	} else if D < 0 {
		fmt.Println("Уравнение не имеет корней \n D < 0 \n D = " + fmt.Sprint(D))
	}
	fmt.Scan(&a)

	// Команда ддля оптимизированной сборки
	//go build -ldflags "-s -w"
}
