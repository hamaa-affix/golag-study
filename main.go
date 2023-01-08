package main

import (
	"errors"
	"fmt"
	// "log"
)
var slice = []string{"Golang", "Ruby", "Javascript", "Python"}

func main() {
    // コンソールにメッセージが出力される
	// hoge()
	
	// if err := isString(32); err == nil {
	// 	fmt.Println("文字列です")
	// } else {
	// 	// log.Fatal("not string")
	// 	log.Fatal(err)
	// }
	
	for _, value := range slice {
		fmt.Println( value)
	}


}

func isString(arg interface{}) error {
	isString := interface{}(arg)

	if _, ok := isString.(string); !ok {
		return errors.New("errorですよ")
	}

	return nil
}
