package main

import (
	"fmt"
)

func sum(num1 float32, num2 float32) float32 {
	return num1 + num2
}
func min(num1 float32, num2 float32) float32 {
	return num1 - num2
}
func mul(num1 float32, num2 float32) float32 {
	return num1 * num2
}
func div(num1 float32, num2 float32) float32 {
	if num1 != 0 && num2 != 0 {
		return num1 / num2
	} else {
		fmt.Println("Error. Wrong input info")
		return 0
	} // here must be error output else
}

func main() {
	var user_num1 int32
	var user_num2 int32
	var operator string
	fmt.Println("input 1st num: ")
	fmt.Scan(&user_num1)
	fmt.Println("input 2nd num: ")
	fmt.Scan(&user_num2)
	fmt.Println("input the operator: ")
	fmt.Scan(&operator)
	if operator == "+" {
		fmt.Println(sum(float32(user_num1), float32(user_num2)))
	}
	if operator == "-" {
		fmt.Println(min(float32(user_num1), float32(user_num2)))
	}
	if operator == "*" {
		fmt.Println(mul(float32(user_num1), float32(user_num2)))
	}
	if operator == "/" {
		fmt.Println(div(float32(user_num1), float32(user_num2)))
	}
}
