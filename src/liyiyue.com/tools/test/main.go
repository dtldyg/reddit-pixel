package main

import "fmt"

func main() {
	fmt.Println("input")
	var name string
	fmt.Scanln(&name)
	fmt.Println("hello", name)
}
