package main

import (
	"fmt"
	"math/rand"
)

func ex2() {
	slice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		randInt := rand.Intn(100)
		if randInt%2 == 0 && randInt%3 == 0 {
			fmt.Println("Six!")
		}
		if randInt%2 == 0 {
			fmt.Println("Two!")
		}
		if randInt%3 == 0 {
			fmt.Println("Three!")
		}
		fmt.Println("Never mind")
		slice = append(slice, randInt)
	}
	fmt.Println(slice)
	return
}
