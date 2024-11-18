package trie

import (
	"fmt"
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
	trie := New(&Options{IgnoreDiacritics: true})

	trie.ParseText(thekoran, StandardReplacer)
	// trie.Stats()
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
		// tail = "..."
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

func TestTrieJabberwock(t *testing.T) {
	testWords := []string{
		"fox",
		"twice",
		"jabber",
		"jub",
		"sword",
		"vorpal",
		"brillig",
		"sought",
	}
	trie := New(&Options{IgnoreDiacritics: true})
	trie.ParseFile("the-jabberwocky.txt", StandardReplacer)
	// trie.Stats()

	for _, word := range testWords {
		total, ln := trie.Search(word)
		t.Logf("Found %d words containing %s ", total, word)
		if len(ln) > 0 {
			t.Logf("at line %s", printLineNumbers(ln))
		}
	}
}
func TestTrieElefantsChild(t *testing.T) {
	testWords := []string{
		"fox",
		"twice",
		"Crocodile",
		"Crocodiles",
		"pedr",
		"Beloved",
	}
	trie := New(&Options{IgnoreDiacritics: true})
	trie.ParseFile("the-elefants-child.txt", StandardReplacer)
	// trie.Stats()

	for _, word := range testWords {
		total, ln := trie.Search(word)
		t.Logf("Found %d words containing %s ", total, word)
		if len(ln) > 0 {
			t.Logf("at line %s", printLineNumbers(ln))
		}
	}
}

func TestTrieMuchAdo(t *testing.T) {
	trie := New(&Options{IgnoreDiacritics: true})

	trie.ParseText(muchado, StandardReplacer)
	// trie.Stats()
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

type Device struct {
	ID string `json:"id"`
}

func TestReplacer(t *testing.T) {
	// 	str := `‘Hljóðs bið ek allar kindir,
	// meiri ok minni, mǫgu Heimdallar! Vildu at ek, Valfǫðr, vel fyrtelja forn spjǫll fira, þau er fremst um man.
	// `
	str := "Übeltäter übergibt 'Ärzten' öfters äußerst ätzende Öle."
	// var Tr = func() transform.Transformer {
	// 	return transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	// }
	// trans = func(word string) string {
	// 	return norm.NFD.String(runes.Remove(runes.In(unicode.Mn)).String(cases.Lower(language.English).String(word)))
	// }
	a, _, err := transform.String(ToLower(), str)
	if err != nil {
		t.Error(err)
	}
	b, _, err := transform.String(RemoveDiacritics(), str)
	if err != nil {
		t.Error(err)
	}

	t.Log("str", str)
	t.Log("a", a)
	t.Log("b", b)
}

func TestParseItem(t *testing.T) {
	trie := New(&Options{IgnoreDiacritics: false, MaxWordLength: 4})

	trie.ParseItem("123-123-123", StandardReplacer, &Item{
		Description: "momerath device id",
		Path:        "/devices/213-123-123",
	})
	trie.ParseItem("onetwothree.local", StandardReplacer, &Item{
		Description: "momerath device hostname",
		Path:        "/devices/213-123-123",
	})
	trie.ParseItem("powell", StandardReplacer, &Item{
		Description: "phil surname",
		Path:        "/people/phil",
	})
	trie.ParseItem("phil", StandardReplacer, &Item{
		Description: "phil nickname",
		Path:        "/people/phil",
	})
	trie.ParseItem("philip", StandardReplacer, &Item{
		Description: "phil first name",
		Path:        "/people/phil",
	})
	trie.ParseItem("phil.local", StandardReplacer, &Item{
		Description: "phil hostname",
		Path:        "/devices/phil-123-456",
	})
	trie.ParseItem("Übeltäter übergibt 'Ärzten' öfters äußerst ätzende Öle.", StandardReplacer, &Item{
		Description: "german phrase",
		Path:        "/misc/words",
	})
	trie.Stats()
	testWords := []string{
		"123",
		"phil",
		"one",
		"two",
		"three",
		"pow",
		"öfters",
		"ofters",
		"kømr",
		"annarr",
		"vindheim",
	}

	for _, word := range testWords {
		items := trie.SearchItem(word)
		if items != nil {
			t.Logf("Found %d items containing %s ", len(items), word)
			for _, item := range items {
				t.Logf("path %s, description %s", item.Path, item.Description)
			}
		} else {
			t.Logf("No items containing %s found", word)

		}
		// t.Logf("Found %d words containing %s ", total, word)
		// if len(ln) > 0 {
		// 	t.Logf("at line %s", printLineNumbers(ln))
		// }
	}
}
func TestTrieEdda(t *testing.T) {
	trie := New(&Options{IgnoreDiacritics: false})

	trie.ParseText(voluspa, StandardReplacer)
	// trie.Stats()
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
	testWords := []string{
		"fox",
		"bricks",
		"beetle",
		"socks",
		"knox",
		"broom",
		"puddle",
		"noodle",
	}
	trie := New(&Options{IgnoreDiacritics: true})

	trie.ParseFile("fox-in-socks.txt", StandardReplacer)
	// trie.Stats()

	for _, word := range testWords {
		total, ln := trie.Search(word)
		t.Logf("Found %d words containing %s ", total, word)
		if len(ln) > 0 {
			t.Logf("at line %s", printLineNumbers(ln))
		}
	}
}
func TestTrieUUIDs(t *testing.T) {
	testWords := []string{
		"37bf6ca3-cd28-59d4-bf69-80f51e22f407",
		"0bcfa996-4f5e-534c-b6d3",
		"37bf6ca3",
		"socks",
		"knox",
		"broom",
		"puddle",
		"noodle",
	}
	trie := New(&Options{IgnoreDiacritics: true})

	trie.ParseText("37bf6ca3-cd28-59d4-bf69-80f51e22f407", TechnicalReplacer)
	trie.ParseText("0bcfa996-4f5e-534c-b6d3-9e5cd001032f", TechnicalReplacer)
	// trie.Stats()

	for _, word := range testWords {
		total, ln := trie.Search(word)
		if len(ln) > 0 {
			t.Logf("Found %d words containing %s at line %s", total, word, printLineNumbers(ln))
		} else {
			t.Logf("Found no words containing %s ", word)
		}
	}
}

func BenchmarkMuchAdo(b *testing.B) {
	trie := New(&Options{IgnoreDiacritics: true})
	// for range b.N {
	trie.ParseText(muchado, StandardReplacer)
	// trie.Stats()
	// }
}
func BenchmarkKoran(b *testing.B) {
	trie := New(&Options{IgnoreDiacritics: true})
	// for range b.N {
	trie.ParseText(thekoran, StandardReplacer)
	// trie.Stats()
	// }
}
func BenchmarkEdda(b *testing.B) {
	trie := New(&Options{IgnoreDiacritics: false})
	// for range b.N {
	trie.ParseText(voluspa, StandardReplacer)
	// trie.Stats()
	// }
}

func TestTrieSize(t *testing.T) {
	trie := New(&Options{IgnoreDiacritics: true})
	// for range b.N {
	trie.ParseText(muchado, StandardReplacer)

	// trie.Stats()

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
	trie := New(&Options{IgnoreDiacritics: true})

	// Test empty string insertion
	trie.ParseText("", TechnicalReplacer)
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
	trie := New(&Options{IgnoreDiacritics: true})

	specialWords := []string{"hello!", "hello?", "hello-world", "hello_world"}
	for _, word := range specialWords {
		trie.ParseText(word, TechnicalReplacer)
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
	transf := transform.Chain(cases.Lower(language.English))
	word := "BigWord"

	for range b.N {
		word, _, _ = transform.String(transf, word)
	}
}
func BenchmarkBoth1(b *testing.B) {
	transf := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC)
	word := "þá mun Friggjar falla Angantýr."

	for range b.N {
		word, _, _ = transform.String(transf, strings.ToLower(word))
	}
}
func BenchmarkBoth2(b *testing.B) {
	transf := transform.Chain(cases.Lower(language.English), norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC)
	word := "þá mun Friggjar falla Angantýr."

	for range b.N {
		word, _, _ = transform.String(transf, word)
	}
}
func TestBoth2(t *testing.T) {
	tr := transform.Chain(cases.Lower(language.English), norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	word := "‘Þá kømr inn mikli mǫgr Sigfǫður, Víðarr, vega at valdýri; öäü"
	t.Log(word)

	word, _, _ = transform.String(tr, word)
	t.Log(word)
}
