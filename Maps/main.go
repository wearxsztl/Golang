package main

import "fmt"

func main() {
	//var money map[string]int = map[string]int{
	money := map[string]int{
		"dollars": 1000,
		"euros":   2000,
		"rubles":  10000,
		"apples":  6,
	}

	// money["dollars"] = 5000
	// delete(money, "apples")
	// fmt.Println(money)
	// fmt.Println(money["dollars"])

	el, status := money["dollars"] // el ==1000, status == true
	//or el, _ := money["dollars"]
	fmt.Println(el, status)

}
