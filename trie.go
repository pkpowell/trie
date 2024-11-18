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
	"golang.org/x/text/unicode/norm"
)

type Options struct {
	IgnoreDiacritics bool
	MaxWordLength    int
	Replacer         *strings.Replacer
}

type Item struct {
	// Category    string `json:"category,omitempty"`
	Path        string `json:"path,omitempty"`
	Description string `json:"description,omitempty"`
}

type Meta struct {
	Word  string  `json:"word,omitempty"`
	Items []*Item `json:"items,omitempty"`
}

type Node struct {
	Children    Children                 `json:"children,omitempty"`
	IsEnd       bool                     `json:"isEnd"`
	Meta        *Meta                    `json:"meta,omitempty"`
	Count       int                      `json:"count,omitempty"`
	WordCount   int                      `json:"wordCount,omitempty"`
	Lines       Lines                    `json:"lines,omitempty"`
	options     *Options                 `json:"-"`
	mtx         *sync.RWMutex            `json:"-"`
	transformer func(word string) string `json:"-"`
}

type Lines map[int]int

type Children map[string]*Node

var newNode = New
var Defaults = &Options{
	IgnoreDiacritics: true,
	MaxWordLength:    3,
}
var Opts = Defaults

// var transf transform.Transformer
var trans func(word string) string

// New initializes a new Trie
func New(opts *Options) *Node {
	if opts == nil {
		opts = Defaults
	}
	if opts.IgnoreDiacritics {
		trans = func(word string) string {
			return norm.NFD.String(runes.Remove(runes.In(unicode.Mn)).String(cases.Lower(language.English).String(word)))
		}
	} else {
		trans = func(word string) string {
			return norm.NFD.String(cases.Lower(language.English).String(word))
		}
	}

	return &Node{
		Children:  make(Children),
		IsEnd:     false,
		Count:     0,
		WordCount: 0,
		Lines:     make(Lines),
		Meta: &Meta{
			Items: make([]*Item, 0),
		},
		options:     opts,
		mtx:         new(sync.RWMutex),
		transformer: trans,
	}
}

func (root *Node) ParseFile(filename string, replacer *strings.Replacer) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	root.ParseText(string(b), replacer)
}

// ParseItem removes formatting and special characters before adding words to trie
func (root *Node) ParseItem(text string, replacer *strings.Replacer, item *Item) {
	switch true {
	case root == nil:
		fmt.Printf("root is nil")
		return
	case len(text) <= Opts.MaxWordLength:
		return
		// return
	case replacer == nil:
		replacer = StandardReplacer
	}

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		words := strings.Split(replacer.Replace(line), " ")
		for _, word := range words {
			if len(word) < 3 {
				continue
			}

			word := norm.NFC.String(trans(word))

			for i := range len(word) {
				root.update(word[i:], item)
			}

			root.WordCount += len(words)
		}
	}
}

// ParseText removes formatting and special characters before adding words to trie
func (root *Node) ParseText(text string, replacer *strings.Replacer) {

	switch true {
	case root == nil:
		fmt.Printf("root is nil")
		return
	case len(text) <= Opts.MaxWordLength:
		return
		// return
	case replacer == nil:
		replacer = StandardReplacer
	}

	lines := strings.Split(text, "\n")
	for num, line := range lines {
		words := strings.Split(replacer.Replace(line), " ")
		for _, word := range words {
			if len(word) < 2 {
				continue
			}

			word := root.transformer(word)

			for i := range len(word) {
				root.updateWithLines(word[i:], num)
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
func (root *Node) update(word string, item *Item) {
	root.mtx.Lock()
	defer root.mtx.Unlock()

	current := root

	for _, letter := range word {
		_, ok := current.Children[string(letter)]
		if !ok {
			current.Children[string(letter)] = newNode(&Options{IgnoreDiacritics: root.options.IgnoreDiacritics})
		}
		current = current.Children[string(letter)]

	}
	current.Meta = &Meta{
		Word:  word,
		Items: append(current.Meta.Items, item),
	}
	current.IsEnd = true
}

// Update updateWithLines a word in the trie and adds line information
func (root *Node) updateWithLines(word string, num int) {
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
	// var err error
	var node = root

	root.mtx.RLock()
	defer root.mtx.RUnlock()

	// fmt.Println("transf", transf)

	word = root.transformer(word)
	fmt.Println("searching word", word)
	// if err != nil {
	// 	panic(err)
	// }

	for _, letter = range word {
		_, ok = node.Children[string(letter)]
		if !ok {
			return 0, nil
		}
		node = node.Children[string(letter)]
	}
	fmt.Println("found word", node.Meta.Word)
	return node.Count, node.Lines
	// return current.isEnd
}

// search for exact word
func (root *Node) SearchItem(word string) []*Item {
	var letter rune
	var ok bool
	// var err error

	root.mtx.RLock()
	defer root.mtx.RUnlock()

	var node = root
	word = strings.ToLower(strings.Trim(word, " "))

	for _, letter = range word {
		_, ok = node.Children[string(letter)]
		if !ok {
			return nil
		}
		node = node.Children[string(letter)]
	}

	if len(node.Meta.Items) == 0 {
		var extra = []*Item{}
		var more func(n *Node)
		more = func(n *Node) {
			for _, n := range n.Children {
				if n.Meta.Items != nil {
					extra = append(extra, n.Meta.Items...)
					more(n)
				}
			}
		}
		more(node)

		// fmt.Println("found extra words", extra[0].Description)
		return extra
	}
	// fmt.Println("found word", node.Meta.Items)
	return node.Meta.Items
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

	data, err := json.Marshal(root)
	if err != nil {
		fmt.Println("json.Marshal error", err)
		return
	}

	fmt.Printf(("Memory: %d bytes\n"), len(data))
	// fmt.Printf(("%d letters\n"), letters)
	fmt.Printf(("%d words\n"), root.WordCount)
}
