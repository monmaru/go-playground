package tests

import "testing"

// go testでの実行対象にするため、必ずTestから始まる名前にする
func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatal("sum(1, 2) shoud be 3, but does'nt match")
	}
}
