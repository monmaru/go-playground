package playground

import (
	"fmt"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	t := tokenizer.New()
	tokens := t.Tokenize("寿司が食べたい。")
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			fmt.Printf("%s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)
	}
	/*
		BOS
		寿司	名詞,一般,*,*,*,*,寿司,スシ,スシ
		が	助詞,格助詞,一般,*,*,*,が,ガ,ガ
		食べ	動詞,自立,*,*,一段,連用形,食べる,タベ,タベ
		たい	助動詞,*,*,*,特殊・タイ,基本形,たい,タイ,タイ
		。	記号,句点,*,*,*,*,。,。,。
		EOS
	*/
}
