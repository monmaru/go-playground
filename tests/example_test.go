package tests

import "fmt"

/*
Example functionsはExampleから始まる名前で定義し、
出力を// Output:から始まるコメントで書くことで、標準出力の内容をテストできる。
もし、Outputのコメントがない場合にはコンパイルのみが走る。
*/
func ExampleHello() {
	fmt.Println("Hello")
	// Output: Hello
}

/*
Unordered outputoを使うと、順不同な結果に対してもマッチさせることができる。
例えば、mapをイテレートすると順不同に結果が返ってくるので、そういったケースで使用する。
*/
func ExampleUnordered() {
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	// Unordered output:
	// 2
	// 3
	// 1
}
