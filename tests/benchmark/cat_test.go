package benchmark

import (
	"fmt"
	"testing"
)

// seedはベンチマーク用のトークンを作っている
// 長さを受取り、指定された長さの文字列のスライスを生成する
func seed(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, "a")
	}
	return s
}

// benchはベンチマーク用のヘルパーです
// テストしたい文字列の組み合わせ長と、文字列結合のための手続きを渡します
func bench(b *testing.B, n int, f func(...string) string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(seed(n)...)
	}
}

func BenchmarkCat3(b *testing.B) {
	bench(b, 3, cat)
}

func BenchmarkBuf3(b *testing.B) {
	bench(b, 3, buf)
}

func BenchmarkJoin3(b *testing.B) {
	bench(b, 3, join)
}

func BenchmarkCat100(b *testing.B) {
	bench(b, 100, cat)
}

func BenchmarkBuf100(b *testing.B) {
	bench(b, 100, buf)
}

func BenchmarkJoin100(b *testing.B) {
	bench(b, 100, join)
}

func BenchmarkCat10000(b *testing.B) {
	bench(b, 10000, cat)
}

func BenchmarkBuf10000(b *testing.B) {
	bench(b, 10000, buf)
}

func BenchmarkJoin1000(b *testing.B) {
	bench(b, 10000, join)
}

/*
Go1.7からサブベンチマークがサポートされた
サブベンチマークを使うと、1つのBenchmarkXXX手続きの中に複数のベンチマークを記述できる
*/
func BenchmarkConcatenate(b *testing.B) {
	benchCases := []struct {
		name string
		n    int
		f    func(...string) string
	}{
		{"Cat", 3, cat},
		{"Buf", 3, buf},
		{"Join", 3, join},
		{"Cat", 100, cat},
		{"Buf", 100, buf},
		{"Join", 100, join},
		{"Cat", 10000, cat},
		{"Buf", 10000, buf},
		{"Join", 10000, join},
	}

	for _, c := range benchCases {
		b.Run(
			fmt.Sprintf("%s%d", c.name, c.n),
			func(b *testing.B) {
				bench(b, c.n, c.f)
			})
	}
}
