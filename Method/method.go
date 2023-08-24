package main

import "fmt"

type account struct {
	balance int
}

// 일반 함수
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 메서드
func (a *account) withdrawFunc(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100} // account구조체의 포인터 변수 생성

	withdrawFunc(a, 30)
	a.withdrawFunc(30)
	fmt.Printf("%d \n", a.balance)
}
