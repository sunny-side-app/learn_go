package main

import "fmt"

func ex3() {

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}
