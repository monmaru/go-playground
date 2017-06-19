package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Protocol Version: %s\n", resp.Proto)
}

// GODEBUG=http2client=1 go run main.go
// HTTP/2関連のデバッグログを有効にする
