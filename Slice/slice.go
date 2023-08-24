package main

import "fmt"

func main() {
	// var slice1 = []int{1, 2, 3, 4, 5}
	// // var slice2 = make([]int, 3)	// 길이 3이고 원소 개수도 3개
	// var slice2 = make([]int, 3, 5)	// 길이가 5고 원소 개수는 3개

	// var slice = []int{1, 2, 3}
	// for i := 0; i < len(slice); i++ {
	// 	slice[i] += 10
	// 	fmt.Println(slice[i])
	// }
	// for i, v := range slice {
	// 	slice[i] = v * 2
	// 	fmt.Println(slice[i])
	// }

	// var slice = []int{1, 2, 3}
	// fmt.Println(slice, len(slice))
	// slice = append(slice, 4)
	// fmt.Println(slice, len(slice))
	// slice = append(slice, 5, 6, 7, 8)
	// fmt.Println(slice, len(slice))

	var slice1 = make([]int, 3, 3)
	slice2 := append(slice1, 4)
	slice3 := [7]int{10, 20, 30, 40, 50}
	slice1[0] = 100
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3)
}
