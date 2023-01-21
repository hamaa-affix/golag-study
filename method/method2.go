package method

import (
	"fmt"
)

/*
	User型を作成して、作成しnameを取得するgetterを作成してください。
	user : {
		name: string
		age: int
	}
*/

//user型
type User struct {
	name string
	age int
}
//getter methodの定義
func (u *User)getter()string {
	return u.name
}

func (u *User)setter(name string) *User {
	u.name = name
	return u
}

func Method2() {
	user := User{name: "tmp1",age: 90}
	fmt.Println(user.getter())
	user.setter("tmp2")
	fmt.Println(user)
}
