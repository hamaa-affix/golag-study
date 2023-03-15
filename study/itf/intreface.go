package itf

import (
	"fmt"
)

// user profile getterメソッドを定義しているinterfaceを定義してください(NameGetterなど)
type userProfile interface {
	NameGetter() string
}
// User型を定義してください
//user型
type User struct {
	name string
	age int
}
// 上記で作成したinterfaceをUser型で実装してください
func (u User) NameGetter() string{
	return u.name
}
//methodを基底とするuser定義型を作成して、interfaceを実装してください

// func (up userProfile)f() {

// }

func Itf() {
	// 定義したinterfaceを元に初期化してください
	var user userProfile
	// 初期化したinterfaceにUserが代入できるかを確認してみてください
	user = User{
		name:"hoge",
		age:20,
	}
	fmt.Println(user.NameGetter())
}
