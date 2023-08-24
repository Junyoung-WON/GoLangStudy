package main

import "fmt"

// 구조체 선언
// type House struct {
// 	Address string
// 	Size    int
// 	Price   float64
// 	Type    string
// }

// 중첩 구조체
type User struct {
	Name string
	ID   string
	Age  int
}
type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	/*
		구조체 생성 및 초기화
	*/
	// var house House
	// house.Address = "대구광역시 북구 ..."
	// house.Size = 32
	// house.Price = 2.8
	// house.Type = "아파트"
	// var house House = House{
	// 	"대구광역시 북구 ...",
	// 	32,
	// 	2.8,
	// 	"아파트"}

	// fmt.Println("주소:", house.Address)
	// fmt.Printf("크기: %d평\n", house.Size)
	// fmt.Printf("가격: %.2f억 원\n", house.Price)
	// fmt.Println("타입:", house.Type)

	user := User{"원준영", "WON", 25}
	vip := VIPUser{user, 3, 1000}
	fmt.Printf("유저: %s, ID: %s, 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s, ID: %s, 나이: %d, VIP레벨: %d, 가격: %d만원", vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)
}
