package trie

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// var sentence = "Sphinx of black quartz, judge my vow. Jackdaws love my big sphinx of quartz. Pack my box with five dozen liquor jugs. The quick onyx goblin jumps over the lazy dwarf. The quick brown fox jumps over the lazy dog."

func TestTrieKoran(t *testing.T) {
	trie := NewTrie()

	trie.Parse(thekoran)
	trie.Stats()
	testWords := []string{
		"fox",
		"Muhammad",
		"XCII",
		"twice",
		"unlce",
		"pedr",
		"Fatima",
		"Hafsa",
		"Maryam",
		"uncl",
		"uncle",
		"per",
		"valiant",
		"too",
		"hello",
		"help",
		"helped",
		"helper",
		"hell",
		"the",
		"an",
		"abcd",
	}

	for _, word := range testWords {
		total, ln := trie.Search(word)
		t.Logf("Found %d words containing %s ", total, word)
		if len(ln) > 0 {
			t.Logf("at line %s", printLineNumbers(ln))
		}
	}
}

func printLineNumbers(lineNumbers Lines) (ln string) {
	tail := ""
	limit := len(lineNumbers)
	keys := make([]int, len(lineNumbers))
	i := 0
	for k := range lineNumbers {
		keys[i] = k
		i++
	}
	slices.Sort(keys)
	if len(keys) >= limit {
		keys = keys[:min(len(keys), limit)]
		tail = "..."
	}

	var lns []string
	for _, i := range keys {
		if lineNumbers[i] > 1 {
			lns = append(lns, fmt.Sprintf("%d (%d)", i, lineNumbers[i]))
		} else {
			lns = append(lns, fmt.Sprintf("%d", i))
		}
	}

	return strings.Join(lns, ", ") + tail
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func TestTrieElefantsChild(t *testing.T) {
	b, err := os.ReadFile("the-elefants-child.txt") // just pass the file name
	if err != nil {
		panic(err)
	}

	testWords := []string{
		"fox",
		"twice",
		"Crocodile",
		"Crocodiles",
		"pedr",
		"Beloved",
	}
	trie := NewTrie()
	str := string(b) // convert content to a 'string'
	trie.Parse(str)
	trie.Stats()

	for _, word := range testWords {
		total, ln := trie.Search(word)
		t.Logf("Found %d words containing %s ", total, word)
		if len(ln) > 0 {
			t.Logf("at line %s", printLineNumbers(ln))
		}
	}
}

func TestTrieMuchAdo(t *testing.T) {
	trie := NewTrie()

	trie.Parse(muchado)
	trie.Stats()
	testWords := []string{
		"fox",
		"twice",
		"unlce",
		"pedr",
		"LEONATO",
		"ONATO",
		"leon",
		"uncl",
		"uncle",
		"per",
		"valiant",
		"too",
		"hello",
		"help",
		"helper",
		"helped",
		"hell",
		"the",
		"an",
		"Arragon",
	}

	for _, word := range testWords {
		total, ln := trie.Search(word)
		t.Logf("Found %d words containing %s ", total, word)
		if len(ln) > 0 {
			t.Logf("at line %s", printLineNumbers(ln))
		}
	}
}

func TestTrieEdda(t *testing.T) {
	trie := NewTrie()

	trie.Parse(voluspa)
	trie.Stats()
	testWords := []string{
		"Vindálfr",
		"Gandálfr",
		"gandálf",
		"Gandalf",
		"dverga",
		"kømr",
		"annarr",
		"vindheim",
	}

	for _, word := range testWords {
		total, ln := trie.Search(word)
		t.Logf("Found %d words containing %s ", total, word)
		if len(ln) > 0 {
			t.Logf("at line %s", printLineNumbers(ln))
		}
	}
}
func TestTrieFoxInSocks(t *testing.T) {
	b, err := os.ReadFile("fox-in-socks.txt") // just pass the file name
	if err != nil {
		panic(err)
	}

	testWords := []string{
		"fox",
		"bricks",
		"beetle",
		"socks",
		"knox",
		"broom",
		"puddle",
	}
	trie := NewTrie()
	str := string(b) // convert content to a 'string'
	trie.Parse(str)
	trie.Stats()

	for _, word := range testWords {
		total, ln := trie.Search(word)
		t.Logf("Found %d words containing %s ", total, word)
		if len(ln) > 0 {
			t.Logf("at line %s", printLineNumbers(ln))
		}
	}
}

func BenchmarkMuchAdo(b *testing.B) {
	trie := NewTrie()
	// for range b.N {
	trie.Parse(muchado)
	trie.Stats()
	// }
}
func BenchmarkKoran(b *testing.B) {
	trie := NewTrie()
	// for range b.N {
	trie.Parse(thekoran)
	trie.Stats()
	// }
}
func BenchmarkEdda(b *testing.B) {
	trie := NewTrie()
	// for range b.N {
	trie.Parse(voluspa)
	trie.Stats()
	// }
}

func TestTrieSize(t *testing.T) {
	trie := NewTrie()
	// for range b.N {
	trie.Parse(muchado)

	trie.Stats()

}

// func TestTrieBasicOperations(t *testing.T) {
// 	trie := NewTrie()

// 	// Test adding a word
// 	trie.Add("hello")
// 	if trie.Search("hello") == 0 {
// 		t.Error("Expected to find 'hello' in trie")
// 	}

// 	// Test non-existent word
// 	if trie.Search("world") == 0 {
// 		t.Error("Expected 'world' to not be in trie")
// 	}

// 	// Test prefix search
// 	trie.Add("help")
// 	trie.Add("hell")
// 	if trie.Search("hel") == 0 {
// 		t.Error("Expected to find prefix 'hel' in trie")
// 	}

// 	// Test case sensitivity
// 	if trie.Search("Hello") == 0 {
// 		t.Error("Trie should be case sensitive")
// 	}
// }

func TestTrieEmptyString(t *testing.T) {
	trie := NewTrie()

	// Test empty string insertion
	trie.Parse("")
	// if !trie.Search("") {
	// 	t.Error("Expected to find empty string in trie")
	// }
	// if !trie.StartsWith("") {
	// 	t.Error("Expected empty string to be a valid prefix")
	// }
}

// func TestTrieOverlappingWords(t *testing.T) {
// 	trie := NewTrie()

// 	words := []string{"a", "ab", "abc", "abcd"}
// 	for _, word := range words {
// 		trie.Parse(word)
// 	}

// 	for _, word := range words {
// 		if len(trie.Search(word)) == 0 {
// 			t.Errorf("Expected to find '%s' in trie", word)
// 		}
// 		if len(trie.Search(word)) == 0 {
// 			t.Errorf("Expected '%s' substring", word)
// 		}
// 	}
// }

func TestTrieSpecialCharacters(t *testing.T) {
	trie := NewTrie()

	specialWords := []string{"hello!", "hello?", "hello-world", "hello_world"}
	for _, word := range specialWords {
		trie.Parse(word)
		// if !trie.Search(word) {
		// 	t.Errorf("Expected to find '%s' in trie", word)
		// }
	}
}

func BenchmarkToLower(b *testing.B) {
	word := "BigWord"
	for range b.N {
		word = strings.ToLower(word)
	}
}
func BenchmarkCaser(b *testing.B) {
	t = transform.Chain(cases.Lower(language.English))
	word := "BigWord"

	for range b.N {
		word, _, _ = transform.String(t, word)
	}
}
func BenchmarkBoth1(b *testing.B) {
	t = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC)
	word := "þá mun Friggjar falla Angantýr."

	for range b.N {
		word, _, _ = transform.String(t, strings.ToLower(word))
	}
}
func BenchmarkBoth2(b *testing.B) {
	t = transform.Chain(cases.Lower(language.English), norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC)
	word := "þá mun Friggjar falla Angantýr."

	for range b.N {
		word, _, _ = transform.String(t, word)
	}
}
func TestBoth2(t *testing.T) {
	tr := transform.Chain(cases.Lower(language.English), norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	word := "‘Þá kømr inn mikli mǫgr Sigfǫður, Víðarr, vega at valdýri; öäü"
	t.Log(word)

	word, _, _ = transform.String(tr, word)
	t.Log(word)
}
