package main

import "fmt"

func ex1() {
	// NG:長さ５のゼロ値の後ろから追加になってしまう greetings := make([]string, 5)
	greetings := make([]string, 0, 5)
	greetings = append(greetings, "Hello")
	greetings = append(greetings, "Hola")
	greetings = append(greetings, "hello2")
	greetings = append(greetings, "こんにちは")
	greetings = append(greetings, "hello3")

	fmt.Println(greetings)
	oneTwoElemSubSlice := greetings[:2]
	twoThreeFourElemSubSlice := greetings[1:4]
	fourFiveElemSubSlice := greetings[3:]
	fmt.Println(oneTwoElemSubSlice)
	fmt.Println(twoThreeFourElemSubSlice)
	fmt.Println(fourFiveElemSubSlice)

	return
}
