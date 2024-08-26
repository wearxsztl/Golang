package main

import (
	"fmt"
	// "math"
)

func main() {
	//////////////////////////// Выводим 5 раз в терминал
	// for i := 1; i < 6; i++ {
	// 	fmt.Printf("Hello %d\n", i)
	// }

	/////////////////////////////Вывод четных чисел
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	// OR
	// for i := 1; i <= 100; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	////////////////////////////Прерываем цикл
	// for i := 1; i <= 100; i++ {
	// 	if i == 50 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("After cycle")

	///////////////////////////////// Делаем из while for
	// i := 0
	// for i<5{
	// 	fmt.Println(i)
	// 	i++
	// }

	///////////////////////////////////Вывод значений в массиве

	// nums := []int{1, 2, 23, 54}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// for index, element := range nums {
	// 	fmt.Printf("Index: %d Element: %d\n", index, element)
	// }

	/////////////////////////////////////Matrix
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// for _, row := range matrix {
	// 	for _, col := range row {
	// 		fmt.Printf("%d ", col)
	// 	}
	// 	fmt.Println()
	// }
	///////////////////////////////Switch

	// name := "Andrea"

	// switch name {
	// case "Jordan":
	// 	fmt.Println("Jordan Hello")
	// case "Kate":
	// 	fmt.Println("Kate Hello")
	// case "John":
	// 	fmt.Println("Hello John")

	// default:
	// 	fmt.Println("I don't know you")
	// }

	// number := 10
	// switch {
	// case number > 1:
	// 	fmt.Println("Number is greater than 1")
	// 	fallthrough // при выполнении условия позволяет идти дальше в то время как break выходит
	// case number < 11:
	// 	fmt.Println("Number < 11")
	// }

	//////////////////////////////////Опреаторы сравнения логические операторы

	// a := 3
	// b := 10
	// if a > 1 && b > 5 { // && - И, || - ИЛИ,
	// 	fmt.Println("True!")
	// }

	//////////////////////////////////Math

	// var a float64 = 144
	// result := math.Sqrt(a)
	// fmt.Println(result)

	// var a float64 = 25.21354325
	// result := math.Ceil(a) //Округляет к большему. Floor округляет к низкому. Round округляет правильно
	// fmt.Println(result)

	////////////////////////////////Калькулятор

	// fmt.Println("Калькулятор")
	// fmt.Println("Какое действие вы хотите выполнить? (+,-,*,/,%) ")

	// var action string
	// fmt.Scan(&action)

	// fmt.Println("Введите первое число: ")
	// var firstValue int64
	// fmt.Scan(&firstValue)

	// fmt.Println("Введите второе число: ")
	// var secondValue int64
	// fmt.Scan(&secondValue)

	// switch {
	// case action == "+":

	// 	fmt.Printf("Сумма чисел: " + fmt.Sprint(firstValue+secondValue))

	// case action == "-":

	// 	fmt.Printf("Разность чисел: " + fmt.Sprint(firstValue-secondValue))

	// case action == "*":

	// 	fmt.Printf("Выражение чисел: " + fmt.Sprint(firstValue*secondValue))

	// case action == "/":

	// 	fmt.Printf("Деление чисел: " + fmt.Sprint(firstValue/secondValue))

	// case action == "%":

	// 	fmt.Printf("Остаток чисел: " + fmt.Sprint(firstValue%secondValue))
	// default:
	// 	fmt.Println("Действие которое вы ввели не найдено")
	// }

	////////////////////////////Arrays

}
