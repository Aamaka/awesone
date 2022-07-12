package main

import "fmt"

func main() {
	myArr()
}

func myArr() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		println("enter a number: ")
		var number = arr[i]
		fmt.Scanf("%d", &number)

		fmt.Print(number, " ")
	}
}
