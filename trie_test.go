package trie

import (
	"strings"
	"testing"
)

func TestTrieSentence(t *testing.T) {
	trie := NewTrie()
	for _, word := range strings.Split(muchado, " ") {
		trie.Add(word)
	}

	// if !trie.Search("brave") {
	// 	t.Error("Expected to find 'brave' in trie")
	// }
	// if !trie.Search("brother") {
	// 	t.Error("Expected to find 'brother' in trie")
	// }
	// if trie.Contains("dancing") == 0 {
	// 	t.Error("Expected to find 'dancing' in trie")
	// }
	// if trie.Contains("end") == 0 {
	// 	t.Error("Expected to find 'end' in trie")
	// }
	t.Logf("Found %d words containing <dancing> in trie", trie.Contains("dancing"))
	t.Logf("Found %d words containing <pipers> in trie", trie.Contains("pipers"))
	t.Logf("Found %d words containing <flout> in trie", trie.Contains("flout"))
}
func TestTrieBasicOperations(t *testing.T) {
	trie := NewTrie()

	// Test adding a word
	trie.Add("hello")
	if !trie.Search("hello") {
		t.Error("Expected to find 'hello' in trie")
	}

	// Test non-existent word
	if trie.Search("world") {
		t.Error("Expected 'world' to not be in trie")
	}

	// Test prefix search
	trie.Add("help")
	trie.Add("hell")
	if trie.Contains("hel") == 0 {
		t.Error("Expected to find prefix 'hel' in trie")
	}

	// Test case sensitivity
	if trie.Search("Hello") {
		t.Error("Trie should be case sensitive")
	}
}

func TestTrieEmptyString(t *testing.T) {
	trie := NewTrie()

	// Test empty string insertion
	trie.Add("")
	// if !trie.Search("") {
	// 	t.Error("Expected to find empty string in trie")
	// }
	// if !trie.StartsWith("") {
	// 	t.Error("Expected empty string to be a valid prefix")
	// }
}

func TestTrieOverlappingWords(t *testing.T) {
	trie := NewTrie()

	words := []string{"a", "ab", "abc", "abcd"}
	for _, word := range words {
		trie.Add(word)
	}

	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Expected to find '%s' in trie", word)
		}
		if trie.Contains(word) == 0 {
			t.Errorf("Expected '%s' substring", word)
		}
	}
}

func TestTrieSpecialCharacters(t *testing.T) {
	trie := NewTrie()

	specialWords := []string{"hello!", "hello?", "hello-world", "hello_world"}
	for _, word := range specialWords {
		trie.Add(word)
		// if !trie.Search(word) {
		// 	t.Errorf("Expected to find '%s' in trie", word)
		// }
	}
}
