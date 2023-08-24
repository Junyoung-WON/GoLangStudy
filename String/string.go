package main

import "fmt"

func main() {
	// str1 := "Hello\tWorld!\n"
	// str2 := `Hello\tWorld!\n`
	// str3 := `Hello
	// World!
	// wowowow`
	// fmt.Println(str1)
	// fmt.Println(str2)
	// fmt.Println(str3)

	// var char rune = '헿'
	// fmt.Printf("\n%T\n", char)
	// fmt.Println(char)
	// fmt.Printf("%c\n", char)

	str1 := "Hello"
	str2 := "World"

	str1 += " " + str2
	fmt.Println(str1)
}
