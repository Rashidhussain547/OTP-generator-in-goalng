package main

import "fmt"

func main() {
	x := [...]int{10, 20, 30, 40}   
	fmt.Println(len(x))
	fmt.Println("value at index 2 is",x[2])
}