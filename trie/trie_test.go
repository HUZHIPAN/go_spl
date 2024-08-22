package trie

import (
	"fmt"
	"testing"
)

func TestTrie_ContainsFin(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("hello2345")
	trie.Insert("world")
	trie.Insert("abcd")
	trie.Insert("bcd")
	trie.Insert("bcdrddd")
	trie.Insert("bc")
	trie.Insert("c")
	trie.Insert("猪啊")
	trie.Insert("猪啊你")
	trie.Insert("杀猪刀")
	trie.Insert("🔪")
	trie.Insert("杀猪🔪")
	trie.Insert("杀猪🔪")
	trie.Insert("@666")
	trie.Insert("②")

	r := trie.Delete("猪啊")

	fmt.Println(r)

	tests := []struct {
		word string
		want bool
	}{
		{"hell", false},
		{"hello23", true},
		{"worl", false},
		{"heloworldheload", true},
		{"helloo", true},
		{"", false},
		{"fhellox", true},
		{"@66②法大师傅", true},
		{"@", false},
		{"🔪", true},
		{"一把杀猪刀啊", true},
		{"一个大猪", false},
		{"一个大猪啊cc", false},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := trie.Contains(tt.word); got != tt.want {
				t.Errorf("Contains(%v) = %v, want %v", tt.word, got, tt.want)
			}
		})
		t.Run(tt.word, func(t *testing.T) {
			if got := trie.Contains(tt.word); got != tt.want {
				t.Errorf("AcContains(%v) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestTrie_Delete(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abcef")
	trie.Insert("feaggd")
	trie.Insert("dsadfg")
	trie.Insert("abc")

	trie.Delete("abcef")
	r := trie.Contains("abc")
	trie.Delete("feaggd")
	trie.Delete("dsadfg")
	trie.Delete("abc")

	r = trie.Contains("abc")

	trie.Insert("hello")
	trie.Insert("hello2345")
	trie.Insert("world")
	trie.Insert("abcd")
	trie.Insert("bcd")
	trie.Insert("bcdrddd")
	trie.Insert("bc")
	trie.Insert("c")
	trie.Insert("猪啊")
	trie.Insert("猪啊你")
	trie.Insert("杀猪刀")
	trie.Insert("🔪")
	trie.Insert("杀猪🔪")
	trie.Insert("杀猪🔪")
	trie.Insert("@666")
	trie.Insert("②")

	//s := trie.size

	r = trie.Delete("猪啊")

	fmt.Println(r)

	tests := []struct {
		word   string
		before bool
		after  bool
	}{
		{"hell", true, false},
		{"我靠", true, false},
		{"3@22", true, false},
		{"back", true, true},
		{"he你您iinill", true, false},
		{"帆帆帆帆a093】【发sd", true, false},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			si := trie.size
			trie.Insert(tt.word)
			if got := trie.Contains(tt.word); got != tt.before {
				t.Errorf("Contains(%v) before = %v, want %v", tt.word, got, tt.before)
			}
			trie.Delete(tt.word)
			if got := trie.Contains(tt.word); got != tt.after {
				t.Errorf("Contains(%v) after = %v, want %v", tt.word, got, tt.after)
			}

			if si != trie.size {
				t.Errorf("%v size error", tt.word)
			}
		})
	}
}
