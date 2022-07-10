package main

import "fmt"

func main() {
	printAscending()
	printDescending()
}

func l() {
	fmt.Println("Enter a number: ")
}

var name int

func printAscending() {
	l()
	fmt.Scanf("%d", &name)
	for i := 0; i < name; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		println()
	}
}

func printDescending() {
	fmt.Scanf("%d", &name)
	for i := 0; i < name; i++ {
		for j := name; j >= i+1; j-- {
			fmt.Print("* ")
		}
		println()
	}

}
