package tests

import (
	"fmt"
	"os"
	"testing"
)

// TestMain はテストの制御フローを作るツール
func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	shutdown()
	os.Exit(exitCode)
}

// setup、shutdownとしているが、とくに関数名は決まりはないので、Goのコードであれば自由に定義できる
func setup() {
	fmt.Println("初期化処理")
}

func shutdown() {
	fmt.Println("終了処理")
}
