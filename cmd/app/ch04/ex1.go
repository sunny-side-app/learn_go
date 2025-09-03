package main

import (
	"fmt"
	"math/rand"
)

func ex1() {
	slice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		randInt := rand.Intn(100)
		slice = append(slice, randInt)
	}
	fmt.Println(slice)
	return
}
