package trie

import (
	"encoding/json"
	"strings"
)

type Node struct {
	Children map[rune]*Node `json:"children"`
	IsEnd    bool           `json:"isEnd"`
	Count    int            `json:"count"`
}

// NewTrie initializes a new Trie
func NewTrie() *Node {
	return &Node{
		Children: make(map[rune]*Node),
		IsEnd:    false,
		Count:    0,
	}
}

var newNode = NewTrie

// Parse removes formatting and special characters before adding words to trie
func (current *Node) Parse(text string) int {
	words := strings.Split(replacer.Replace(text), " ")

	for _, word := range words {
		current.Add(strings.ToLower(word))
	}
	return len(words)
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
	var limit = length

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
func (current *Node) Search(word string) int {
	var letter rune
	var ok bool
	word = strings.ToLower(word)

	for _, letter = range word {
		_, ok = current.Children[letter]
		if !ok {
			return 0
		}
		current = current.Children[letter]
	}
	return current.Count
	// return current.isEnd
}

func (root *Node) Length() int {
	data, err := json.Marshal(root)
	if err != nil {
		panic(err)
	}
	return len(data)
}
