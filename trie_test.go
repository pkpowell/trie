package trie

import (
	"strings"
	"testing"
)

var (
	sentence1 = "Blindtexte werden ebenfalls zur Demonstration der Gestalt verschiedener Schrifttypen und zu Layoutzwecken verwendet. Sie ergeben in der Regel keinen inhaltlichen Sinn. Aufgrund ihrer verbreiteten Funktion als Fülltext für das Layout kommt einer Nicht-Lesbarkeit besondere Bedeutung zu, da die menschliche Wahrnehmung u.a. darauf ausgerichtet ist, bestimmte Muster und Wiederholungen zu erkennen. Ist die Verteilung der Buchstaben und die Länge der \"Worte\" willkürlich, lenkt beispielsweise nichts von der Beurteilung der Wirkung und Lesbarkeit verschiedener Schriftarten (Typografie) sowie der Verteilung des Textes auf der Seite (Layout oder Satzspiegel) ab. Deshalb bestehen Blindtexte meist aus einer mehr oder weniger willkürlichen Folge von Wörtern oder Silben. Wiederholungsmuster können also nicht den Gesamteindruck trüben und Schriftarten so besser miteinander verglichen werden. Dabei ist natürlich von Vorteil, wenn der Blindtext halbwegs realistisch erscheint, damit die Wirkung des Layouts der späteren Publikation nicht beeinträchtigt wird."
	sentence2 = "Als bekanntester Blindtext gilt der Text \"Lorem ipsum\", der seinen Ursprung im 16. Jahrhundert haben soll. Lorem ipsum ist in einer pseudo-lateinischen Sprache verfasst, die ungefähr dem \"natürlichen\" Latein entspricht. In Ihm finden sich eine Reihe realer lateinischer Wörter. Auch dieser Blindtext ist unverständlich gehalten, imitiert jedoch den Rhythmus der meisten europäischen Sprachen in lateinischer Schrift. Der Vorteil des lateinischen Ursprungs und der relativen Sinnlosigkeit von Lorem ipsum ist, dass der Text weder die Aufmerksamkeit des Betrachters auf sich zieht noch von der Gestaltung ablenkt."
)

func TestTrieSentence(t *testing.T) {
	trie := NewTrie()
	for _, word := range strings.Split(sentence2, " ") {
		trie.Add(word)
	}

	if !trie.Search("Ursprung") {
		t.Error("Expected to find 'Ursprung' in trie")
	}
	if !trie.Search("gilt") {
		t.Error("Expected to find 'gilt' in trie")
	}
	if trie.Contains("ert") == 0 {
		t.Error("Expected to find 'ert' in trie")
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

	// Test all prefixes
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Expected to find '%s' in trie", word)
		}
		if trie.Contains(word) == 0 {
			t.Errorf("Expected '%s' to be a valid prefix", word)
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
