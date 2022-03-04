package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	

	var lift = 10
	for lift >0  {
		fmt.Println(lift)
		time.Sleep(time.Second)
		lift--
		if lift == 9 {
			break
		}
	}
	fmt.Println("Liftoff")


	var count = 0 
	for count <10 {
		var num = rand.Intn(20)+1
		fmt.Println(num)
		count ++
	}

	var dt = time.Now()
	fmt.Println("Current date and time in your region is:", dt.Format("2006-01-02 15:04:05 Monday"))
}
