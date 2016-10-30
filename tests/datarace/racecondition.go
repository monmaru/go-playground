package main

import (
	"fmt"
	"sync"
)

/*
Race Detectorは64bitのx86プロセッサを積んだマシンにおいて動作します

- 実行方法
go run -race race.go
*/

func main() {
	c := make(chan bool)
	// 排他制御なしに並行にmapにアクセスすると競合状態が発生します
	// ここではgoroutineからの書き込みとmain goroutineからの書き込みを並行に行って、競合状態を発生させています

	m := make(map[string]string)
	go func() {
		m["1"] = "a" // 1つ目の競合するメモリアクセス
		c <- true
	}()

	m["2"] = "b" // 2つ目の競合するメモリアクセス

	// cはunbuffered channelなのでchannelへのsendが完了するまで待ちます
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func fixed() {
	// mapへの書き込みを排他制御する
	var mux sync.Mutex
	c := make(chan bool)

	m := make(map[string]string)
	go func() {
		mux.Lock()
		m["1"] = "a"
		mux.Unlock()
		c <- true
	}()

	mux.Lock()
	m["2"] = "b"
	mux.Unlock()

	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
