package main

import "fmt"

var global_int int = 32 // 전역변수

func main() {
	var num int = 10            // 기본 형태, 지역변수
	var num2 float64 = 20       // 실수
	var no_init int             // 0으로 자동 초기화
	no_need_type := "wow"       // 타입 지정 x
	var msg string = "Hello YU" // 문자열

	var mult_result = num * (int(num2))

	fmt.Println(num, num2, mult_result, msg)
	fmt.Print(no_need_type, no_init, global_int, "\n")
	fmt.Printf("no_init : %d, no_need_type : %s, global_int : %d", no_init, no_need_type, global_int)

}
