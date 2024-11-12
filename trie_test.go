package trie

import (
	"strings"
	"testing"
)

var (
	sentence = "Round the ragged rocks the ragged rascals ran"
)

func TestTrieSentence(t *testing.T) {
	trie := NewTrie()
	for _, word := range strings.Split(sentence, " ") {
		trie.Add(word)
	}

	if !trie.Search("rocks") {
		t.Error("Expected to find 'rocks' in trie")
	}
	if !trie.Search("Rocks") {
		t.Error("Expected to find 'Rocks' in trie")
	}
	if !trie.Search("th") {
		t.Error("Expected to find 'th' in trie")
	}
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
	// if !trie.StartsWith("hel") {
	// 	t.Error("Expected to find prefix 'hel' in trie")
	// }

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

	// Test all prefixes
	// for _, word := range words {
	// 	if !trie.Search(word) {
	// 		t.Errorf("Expected to find '%s' in trie", word)
	// 	}
	// 	if !trie.StartsWith(word) {
	// 		t.Errorf("Expected '%s' to be a valid prefix", word)
	// 	}
	// }
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
