package main

import "fmt"

func ex2() {
	message := "Hi ✋ and 😉"
	// NG例: fmt.Println(string(message[3]))
	// OK例:
	r := []rune(message) // runeリテラルは1文字のUnicode文字の形式で表現できる
	fmt.Println(string(r[3]))
	return
}
