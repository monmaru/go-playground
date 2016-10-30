package benchmark

import "bytes"
import "strings"

/*
ベンチマークの機能がテスト用パッケージに標準で入っている

- 実行方法
go test -bench .

- 出力結果例
testing: warning: no tests to run
BenchmarkCat3-4       	10000000	       192 ns/op	      54 B/op	       3 allocs/op
BenchmarkBuf3-4       	10000000	       231 ns/op	     163 B/op	       3 allocs/op
BenchmarkJoin3-4      	10000000	       164 ns/op	      54 B/op	       3 allocs/op
BenchmarkCat100-4     	  200000	      6542 ns/op	    7392 B/op	     100 allocs/op
BenchmarkBuf100-4     	 1000000	      2294 ns/op	    2032 B/op	       4 allocs/op
BenchmarkJoin100-4    	 1000000	      1640 ns/op	    1888 B/op	       3 allocs/op
BenchmarkCat10000-4   	     100	  10611378 ns/op	53168049 B/op	   10000 allocs/op
BenchmarkBuf10000-4   	    5000	    210755 ns/op	  208608 B/op	      11 allocs/op
BenchmarkJoin1000-4   	   10000	    157947 ns/op	  184832 B/op	       3 allocs/op
PASS

*/

// catは+=演算子をつかって文字列を結合する
func cat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}
	return r
}

// butはbytes.Bufferをつかって文字列を結合する
func buf(ss ...string) string {
	var b bytes.Buffer
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

// joinはstrings.Joinをつかって文字列を結合する
func join(ss ...string) string {
	return strings.Join(ss, "")
}
