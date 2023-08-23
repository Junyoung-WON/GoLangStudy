package main

import "fmt"

func main() {
	day := "wednesday"

	// break 없어도 알아서 해당 case이후 switch를 빠져나옴
	// 다음 case문 까지 같이 실행하고 싶을 땐 fallthrough 키워드 사용(break 키워드 사용 불필요)
	switch day {
	case "monday", "wednesday":
		fmt.Println("근로 가는 날")
		fallthrough
	case "tuesday", "thursday":
		fmt.Println("수업 가는 날")
	default:
		fmt.Println("공강")
	}

}
