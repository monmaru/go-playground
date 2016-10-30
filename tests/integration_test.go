// +build integration

package tests

import "testing"

func TestSomething(t *testing.T) {
	f, err := OpenSomething("path/to/file")
}

/*
got test
ではTestSomethingは実行されず、
go test -tags=integration
とした場合のみテストが実行される。
*/
