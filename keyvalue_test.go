package playground

import "testing"

func TestSet(t *testing.T) {
	kv := NewKeyValue()
	kv.Set("Key", "value")
	value, ok := kv.Get("Key")

	if !ok {
		t.Error("Setに失敗しました")
	}

	if value != "value" {
		t.Error("戻り値がvalueではありません")
	}
}
