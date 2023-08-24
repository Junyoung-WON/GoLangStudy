package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // WaitGroup 객체 생성
var mutex sync.Mutex  // 패키지 전역 변수 뮤텍스

type Account struct {
	Balance int
}

func PrintHangle() {
	hangles := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hangles {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done() // 작업이 완료됨을 표시
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()                // 뮤텍스 획득
	defer mutex.Unlock()        // defer를 사용하여 해당 함수가 끝나기 전 Unlock()
	account.Balance += 1000     // 1000원 입금
	time.Sleep(time.Nanosecond) // 잠시 쉬고
	account.Balance -= 1000     // 1000원 출금
	wg.Done()
}

func main() {
	// go PrintHangle()
	// go PrintNumbers()
	// time.Sleep(3 * time.Second)

	// wg.Add(10) // 총 작업 개수 설정
	// start := time.Now()
	// for i := 0; i < 10; i++ {
	// 	go SumAtoB(1, 1000000000)
	// }
	// wg.Wait() // 모든 작업이 완료되길 기다림
	// end := time.Since(start)
	// fmt.Println("모든 계산이 완료되었습니다. 걸린 시간:", end)

	// start = time.Now()
	// for i := 0; i < 10; i++ {
	// 	go SumAtoB(1, 1000000000)
	// }
	// time.Sleep(1 * time.Second)
	// end = time.Since(start)
	// fmt.Println("모든 계산이 완료되었습니다. 걸린 시간:", end)

	account := &Account{0} // 0원 통장
	wg.Add(10)             // Go루틴 10개 생성
	for i := 0; i < 10; i++ {
		go DepositAndWithdraw(account)
	}
	wg.Wait()
	fmt.Println(account.Balance)
}
