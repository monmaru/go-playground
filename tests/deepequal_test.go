package tests

import "testing"
import "reflect"

/*
大きいmapやstructを比較する際、forでイテレートしつつ、それぞれの値を比較することもできる
reflect.DeepEqualを使うとより簡単に2つの値が等価であるか否かを比較できる
*/

type T struct {
	x  int
	ss []string
	m  map[string]int
}

func TestStruct(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	t1 := T{
		x:  1,
		ss: []string{"a", "b"},
		m:  m1,
	}

	t2 := T{
		x:  1,
		ss: []string{"a", "b"},
		m:  m1,
	}

	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("want %#v got %#v", t1, t2)
	}
}

/*
次のmapTestのようにテストケースを2つのmapとし等価比較をする場合、
reflect.DeepEqualで結果を比較すると簡単に実装できる
*/

type mapTest struct {
	a, b map[string]int
	eq   bool
}

var mapTests = []mapTest{
	{map[string]int{"a": 1}, map[string]int{"b": 1}, false},
	{map[string]int{"a": 1}, map[string]int{"a": 1}, true},
}

func TestMapTable(t *testing.T) {
	for _, test := range mapTests {
		if r := reflect.DeepEqual(test.a, test.b); r != test.eq {
			t.Errorf("when a = %#v and b = %#v, want %t, got %t", test.a, test.b, r, test.eq)
		}
	}
}
