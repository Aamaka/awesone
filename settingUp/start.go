package main

import "fmt"

func main() {
	fmt.Println("hello, playground")
	add(3)

}

func add(int2 int) {
	i := int2 * 2
	fmt.Println(i)
}
