package slice

import ( 
	"fmt"
)
func Slice() {

	type TMP struct {
		N int
	}
	// var構文を用いてint型のsliceを作成してください。
	var tmp []int
	fmt.Println(tmp)
	// 作成したsliceを元にint値を代入してください
	tmp = append(tmp,1,2)
	fmt.Println(tmp)
	// 型注釈を用いた構文でsliceを作成してください
	tmp2 := []int{1,2}
	fmt.Println(tmp2)
	//make関数を用いてsliceを初期化してください
	tmp3 := make([]int, 5,5)
	fmt.Println(tmp3)
	//appendで要素を追加してください
	tmp3 = append(tmp3, 6)
	fmt.Println(tmp3)
	// N: intを持つstruct型のsliceを作成してください
	tmp4 := []TMP{{N: 6},{N: 3}}
	fmt.Println(tmp4)
}
