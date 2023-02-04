package error

import (
	"errors"
	"fmt"
)

type User struct {
	name string
	age int
	err error
}

func (u *User) NameSetter(newName interface{}) {
	if u.err != nil {
		return
	}
	
	// i := interface{}(newName)
	if _, ok := newName.(string); !ok {
		u.err = errors.New(fmt.Sprintf("エラーだよ:%v",newName))
		return
	}
	
	u.name = newName.(string)

}	

func (u *User) Err() error {
	return u.err
}

func ErrorDemo() {
	user := new(User)
	user.name = "hoge"
	user.age = 100
	fmt.Println(user)
	
	user2 := &User{
		name:"hoge",
		age:20,
	}

	user.NameSetter("f")
	user.NameSetter(10)
	user.NameSetter("a")

	fmt.Println(user2.name)
	if user.Err() != nil {
		fmt.Println(user.err)
	}
	
}

