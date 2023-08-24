package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value += 100
	arg.data[100] += 100
}

func main() {
	// var num int = 500
	// var ptr *int = &num // int 포인터 변수 ptr 선언
	// ptr = &num

	// fmt.Printf("ptr의 값: %p\n", ptr)
	// fmt.Printf("ptr이 가리키는 메모리의 값: %d\n", *ptr)
	// *ptr = 100
	// fmt.Printf("num의 값: %d\n", num)
	// fmt.Printf("ptr이 가리키는 메모리의 값: %d\n", *ptr)

	var data Data

	ChangeData(&data)
	fmt.Printf("value = %d, data[100] = %d\n", data.value, data.data[100])

}
