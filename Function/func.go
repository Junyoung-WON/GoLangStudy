package main

import "fmt"

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	c := Add(3, 10)
	fmt.Println(c)
}
