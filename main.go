package main

import (
	"errors"
	"fmt"
)

func main() {
    // コンソールにメッセージが出力される
	// hoge()
	fmt.Println(Green)
	fmt.Printf("%T\n",Green )

	//3 fizz 5 bazu
	for i := 1; i < 100; i++ {
		if  i % 15 == 0 {
			fmt.Println("fizz bazz")
		} else if i % 5 == 0 {
			fmt.Println("bazz")
		} else if i % 3 == 0  {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func isString(arg string) error {
	isString := interface{}(arg)

	if _, ok := isString.(string); !ok {
		return errors.New("errorですよ")
	}

	return nil
}
