package main

import (
	"fmt"
	"math/rand"
)

func main(){
	min := 10
	max := 30
	var num = rand.Intn(max-min)+max
	fmt.Println(num)
}