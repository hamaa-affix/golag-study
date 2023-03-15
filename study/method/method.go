package method

import (
	"fmt"
)
func Method() {
	//slice
	slice := []int{1, 2, 3, 4, 5}

// int型のスライスを受け取り要素の順番をリバース関数を定義してそのスライスを返却してさい
	s2 := Reverse(slice)
	fmt.Println(s2)
}

// int型のスライスを受け取り要素の順番をリバース関数を定義してそのスライスを返却してさい
func Reverse(s []int) []int {
	reverse := s
	for left, right := 0, len(s) - 1; left < right; left, right = left + 1, right -1 {
		s[left], s[right] = s[right], s[left]
	}

	return reverse
}
