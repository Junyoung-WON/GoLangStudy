package main

import (
	"fmt"
	"os"
)

// func sum(nums ...int) int {
// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	return sum
// }

func main() {
	// fmt.Println(sum(1, 2, 3, 4, 5))
	// fmt.Println(sum(10, 20))
	// fmt.Println(sum())

	// 코드가 작성된 위치와 상관 없이 무조건 해당 함수 종료 전에 실행
	// defer fmt.Println("1. 이 출력은 함수 종료 전에 실행됩니다.")
	// defer fmt.Println("2. 이 출력은 함수 종료 전에 실행됩니다.")
	// defer fmt.Println("3. 이 출력은 함수 종료 전에 실행됩니다.")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	f, err := os.Create("test.txt") // 파일 생성
	if err != nil {                 // 에러 확인
		fmt.Println("Failed to create a file")
		return
	}
	defer fmt.Println("반드시 호출됩니다.") // 지연 수행될 코드
	defer f.Close()                 // 지연 수행될 코드
	defer fmt.Println("파일을 닫았습니다.") // 지연 수행될 코드

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World") // 파일에 텍스트 write
}
