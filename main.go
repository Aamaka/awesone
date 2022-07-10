package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	fmt.Println("Hello World")

	for i := 0; i < 7; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
