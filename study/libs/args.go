package libs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type HttpBody struct {
	code int
	body []byte
	err error
}

func (h *HttpBody) Err() error {
	return h.err
}

func Args() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("引数が少ないです")
	}

	fmt.Printf("hello world\n os.Args %v: %v", args, args[1:])

	// urlのパースチェック
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		log.Fatal(err)
	}

	// http get requestをコール
	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal("receive a message\n $s", err)
	}
	// http.Getで取得したresponseは必ず閉じないといけない
	defer response.Body.Close()

	// io.AllRead()で全てのデータを読み込む
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 標準出力でstatus codeとbodyを出力
	fmt.Printf("HTTP statuscode is %d\nBODY: %s\n", response.StatusCode, string(body))
}
