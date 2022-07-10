package main

import "fmt"

func main() {
	fizz()
}

var userInput int

func input() {
	fmt.Println("Enter number")
}

func fizz() {
	input()
	loop()
}

func loop() {
	fmt.Scanf("%d", &userInput)
	for i := 1; i <= userInput; i++ {
		if i%15 == 0 {
			println("fizzbuzz")
		}
		if i%5 == 0 {
			println("buzz")
		}
		if i%3 == 0 {
			println("fizz")
		} else {
			println(i)
		}
		println()
	}
}
