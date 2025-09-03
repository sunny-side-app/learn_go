package main

import (
	"fmt"
)

func ex4() {
	var total int
	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Println(total)
	return
}
