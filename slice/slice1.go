package slice

import (
	"fmt"
)

/*
    user定義型を作成してslice型に組み込んでください。
	user : {
		name: string
		age: int
	}
*/

type User struct {
	name string
	age int
}



func Slice2() {
	// sliceの作成 sliceの要素は2つ追加してください
	users := []User{{name: "tmp1" ,age: 90},{name: "tmp2" ,age: 90}}
	// for range構文を用いて作成したnameを標準出力してください
	for _, v := range users {
		fmt.Println(v.name)
    }

}
