package main

import (
	"fmt"
	"strings"
)
func main(){
	fmt.Println("you find yourself in the cavern")
	var command = "walk outside"
	var exit = strings.Contains(command,"inside")  
	fmt.Println("you are exit from the cave",exit)
}