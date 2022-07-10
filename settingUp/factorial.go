package main

import "fmt"

func main() {
	fact(6)
}
func fact(number int) {
	result := 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	fmt.Print(result)
}
