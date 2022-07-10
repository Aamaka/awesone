package main

import "fmt"

func main() {
	fmt.Println("Enter number")
	var userInput int
	fmt.Scanf("%d", &userInput)
	multi(userInput)
}

func multi(num int) {
	for i := 0; i < 13; i++ {
		fmt.Println(num, "*", i, "=", num*i)
	}
}
