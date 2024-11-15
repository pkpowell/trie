package trie

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Options struct {
	// CaseSensitive    bool
	IgnoreDiacritics bool
	Replacer         *strings.Replacer
}
type Lines map[int]int

type Node struct {
	Children  map[string]*Node `json:"children"`
	IsEnd     bool             `json:"isEnd"`
	Count     int              `json:"count"`
	WordCount int              `json:"wordCount"`
	Lines     Lines            `json:"line"`
	options   *Options         `json:"-"`
	mtx       *sync.RWMutex    `json:"-"`
}

var trans func(word string) string

// New initializes a new Trie
func New(opts *Options) *Node {
	if opts.IgnoreDiacritics {
		trans = func(word string) string {
			return norm.NFD.String(runes.Remove(runes.In(unicode.Mn)).String(cases.Lower(language.English).String(word)))
		}
		// transf = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC, cases.Lower(language.English))
	} else {
		trans = func(word string) string {
			return norm.NFD.String(cases.Lower(language.English).String(word))
		}
		opts = Defaults
	}

	return &Node{
		Children:  make(map[string]*Node),
		IsEnd:     false,
		Count:     0,
		WordCount: 0,
		Lines:     make(Lines),
		options:   opts,
		mtx:       new(sync.RWMutex),
	}
}

var newNode = New
var Defaults = &Options{
	IgnoreDiacritics: true,
}

var Opts = Defaults

var transf transform.Transformer

func (root *Node) ParseFile(filename string, replacer *strings.Replacer) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	root.ParseText(string(b), replacer)
}

// ParseText removes formatting and special characters before adding words to trie
func (root *Node) ParseText(text string, replacer *strings.Replacer) {
	if root == nil {
		fmt.Printf("root is nil")
		return
	}
	if len(text) < 2 {
		return
	}
	if replacer == nil {
		replacer = StandardReplacer
	}
	// root.mtx.Lock()
	// defer root.mtx.Unlock()

	lines := strings.Split(text, "\n")
	for num, line := range lines {
		words := strings.Split(replacer.Replace(line), " ")
		for _, word := range words {
			if len(word) < 2 {
				continue
			}
			// fmt.Println("word", word)

			// fmt.Printf("word len %d\n", len(word))

			word := norm.NFC.String(trans(word))

			// word, _, err := transform.String(transf, word)
			// if err != nil {
			// 	fmt.Println("transform error", err)
			// 	continue
			// }

			for i := range len(word) {
				root.update(word[i:], num)
			}

		}
		root.WordCount += len(words)
	}
}

// Add adds single exact words
// func (root *Node) Add(word string, num int) {
// 	length := len(word)
// 	if length < 2 {
// 		return
// 	}

// 	// var limit = length // -2

//		// if Opts.IgnoreDiacritics {
//		// 	word, _, err := transform.String(t, word)
//		// 	if err != nil {
//		// 		panic(err)
//		// 	}
//		// 	// fmt.Println(word)
//		// 	// root.update(word)
//		for i := range len(word) {
//			root.update(word[i:], num)
//		}
//		// } else {
//		// 	for i := range len(word) {
//		// 		root.update(word[i:])
//		// 	}
//		// }
//	}
//

// Update updates a word in the trie
func (root *Node) update(word string, num int) {
	root.mtx.Lock()
	defer root.mtx.Unlock()

	current := root

	for _, letter := range word {
		_, ok := current.Children[string(letter)]
		if !ok {
			current.Children[string(letter)] = newNode(&Options{IgnoreDiacritics: root.options.IgnoreDiacritics})
		}
		current = current.Children[string(letter)]
		current.Count++
		current.Lines[num+1]++
	}
	current.IsEnd = true
}

// search for exact word
func (root *Node) Search(word string) (total int, lines Lines) {
	var letter rune
	var ok bool
	var err error
	var node = root

	root.mtx.RLock()
	defer root.mtx.RUnlock()

	word, _, err = transform.String(transf, word)
	if err != nil {
		panic(err)
	}

	for _, letter = range word {
		_, ok = node.Children[string(letter)]
		if !ok {
			return 0, nil
		}
		node = node.Children[string(letter)]
	}
	return node.Count, node.Lines
	// return current.isEnd
}

type Stats struct {
	Words   int
	Letters int
	Memory  int
}

func (root *Node) Stats() {
	root.mtx.RLock()
	defer root.mtx.RUnlock()

	// letters := len(root.Children)
	// for _, node := range root.Children {
	// 	letters += len(node.Children)
	// }
	data, err := json.Marshal(root)
	if err != nil {
		fmt.Println("json.Marshal error", err)
		return
	}

	fmt.Printf(("Memory: %d bytes\n"), len(data))
	// fmt.Printf(("%d letters\n"), letters)
	fmt.Printf(("%d words\n"), root.WordCount)

}
