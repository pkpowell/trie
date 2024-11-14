package trie

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Options struct {
	CaseSensitive    bool
	IgnoreDiacritics bool
}

type Node struct {
	Children  map[rune]*Node `json:"children"`
	IsEnd     bool           `json:"isEnd"`
	Count     int            `json:"count"`
	Tags      []string       `json:"tags"`
	WordCount int            `json:"wordCount"`
}

// NewTrie initializes a new Trie
func NewTrie() *Node {

	return &Node{
		Children:  make(map[rune]*Node),
		IsEnd:     false,
		Count:     0,
		WordCount: 0,
	}
}

var newNode = NewTrie
var Defaults = &Options{
	IgnoreDiacritics: false,
}

var Opts = Defaults
var t = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

// Parse removes formatting and special characters before adding words to trie
func (root *Node) Parse(text string) {
	// if opts != nil {
	// 	Opts = opts
	// }
	// norm.NFC.Transform(text, true)
	// words, _, _ := transform.String(t, text)
	words := strings.Split(replacer.Replace(strings.ToLower(text)), " ")
	for _, word := range words {
		root.Add(strings.ToLower(word))
	}
	root.WordCount = len(words)
}

// Add adds single exact words
func (root *Node) Add(word string) {
	length := len(word)
	if length < 2 {
		return
	}

	var current *Node
	var sub string
	var letter rune
	var ok bool
	var limit = length // -2

	for i := range limit {
		sub = word[i:]
		current = root
		for _, letter = range sub {
			_, ok = current.Children[letter]
			if !ok {
				current.Children[letter] = newNode()
			}
			current = current.Children[letter]
			current.Count++
		}
		current.IsEnd = true
	}
}

// search for exact word
func (root *Node) Search(word string) int {
	var letter rune
	var ok bool
	var err error

	word = strings.ToLower(word)
	// word, _, err = transform.String(t, word)
	if err != nil {
		panic(err)
	}
	// norm.NFC.Transform(word, true)
	node := root

	for _, letter = range word {
		_, ok = node.Children[letter]
		if !ok {
			return 0
		}
		node = node.Children[letter]
	}
	return node.Count
	// return current.isEnd
}

type Stats struct {
	Words   int
	Letters int
	Memory  int
}

func (root *Node) Stats() {
	letters := len(root.Children)
	for _, node := range root.Children {
		letters += len(node.Children)
	}
	data, err := json.Marshal(root)
	if err != nil {
		panic(err)
	}

	fmt.Printf(("Memory: %d bytes\n"), len(data))
	fmt.Printf(("%d letters\n"), letters)
	fmt.Printf(("%d words\n"), root.WordCount)

}
