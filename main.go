package main

import (
	"fmt"
)

func sum(num1 int, num2 int) int {
	return num1 + num2
}
func min(num1 int, num2 int) int {
	return num1 - num2
}
func mul(num1 int, num2 int) int {
	return num1 * num2
}
func div(num1 int, num2 int) int {
	if num1 != 0 && num2 != 0 {
		return num1 / num2
	} else {
		fmt.Errorf(error) // fix error output
		return 0
	}

}

func main() {

	var num1 int = 5
	var num2 int = 10

	fmt.Println(sum(num1, num2))
}
