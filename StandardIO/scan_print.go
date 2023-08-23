package main

import (
	"bufio" // io를 담당하는 패키지
	"fmt"
	"os" // 표준 입출력 등을 가지고 있는 패키지
)

// bufio : 입력 스트림으로부터 한 줄을 읽는 Reader 객체 제공
// NewReader()는 인수로 입력되는 입력 스트림에 대한 Reader 객체를 생성
// 여기서는 표준 입력 스트림을 나타내는 os.Stdin 사용

func main() {
	var a int
	var b int

	stdin := bufio.NewReader(os.Stdin)
	n, err := fmt.Scan(&a, &b)
	if err != nil { // 비정상 입력의 경우
		fmt.Println(err)
		stdin.ReadString('\n') // 표준 입력 스트림 비우기
	} else { // 정상 입력의 경우
		fmt.Println("n:", n, "a:", a, "b:", b)
	}

	// n에는 정상 입력받은 개수
	n, err = fmt.Scan(&a, &b)
	// nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.
	if err != nil { // 비정상 입력의 경우
		fmt.Println(n, err)
	} else { // 정상 입력의 경우
		fmt.Println("n:", n, "a:", a, "b:", b)
	}
}
