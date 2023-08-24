package main

import "fmt"

func main() {
	/*
		배열 선언 및 초기화하기
	*/
	// var arr [5]int	// 초기값 지정 x, 각 0으로 초기화
	// var arr [5]int = int{10, 20, 30, 40, 50}
	// var arr = int{10, 20, 30, 40, 50}
	// arr := int{10, 20, 30, 40, 50}

	/*
		특정 인덱스 값만 초기화
	*/
	// var arr [5]int = [5]int{1: 10, 3: 50}

	/*
		배열 요소 개수 생략
	*/
	// arr := [...]int{0: 5, 10, 2: 30, 50, 50, 7: 1}

	/*
		배열 출력하기
	*/
	// for i := 0; i < len(arr); i++ { // arr 배열 길이만큼 반복
	// 	fmt.Println(arr[i])
	// }

	/*
		range 배열 순회
	*/
	arr := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i, v := range arr { // python처럼 for문에서 range 키워드로 배열 요소를 순회할 수 있음
		fmt.Println(i, v)
	}
	// 인덱스를 쓰고 싶지 않다면 _ 사용!

	/*
		배열 복사
	*/
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for _, v := range arr1 { // python처럼 for문에서 range 키워드로 배열 요소를 순회할 수 있음
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	// 배열 복사 편하게 가능, 단 길이가 같아야함
	arr1 = arr2
	for _, v := range arr1 { // python처럼 for문에서 range 키워드로 배열 요소를 순회할 수 있음
		fmt.Printf("%d ", v)
	}
	fmt.Println()

}
