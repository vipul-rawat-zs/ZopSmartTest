package main

import (
	"fmt"
	day6 "repos/ZopSmartTest/day_6/Functions"
)

func main() {
	// funcx.A()
	// funcx.B()

	// fmt.Println(day6.CheckPallindrome("A man, A plan ,A canal :Panama"))
	// array := [6]int{1, 2, 3}
	// for i := range array {
	// 	array[i] += 2
	// }
	// fmt.Println(array)

	emp := day6.CreateEmployee(0, 1, "Vipul", "Intern", true)
	_, ex := emp.CheckEmpAge()
	fmt.Println(emp, ex)
}
